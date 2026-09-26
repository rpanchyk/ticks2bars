package pipeline

import (
	"context"
	"fmt"
	"sync"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
	"github.com/rpanchyk/ticks2bars/internal/services/broker"
	"github.com/rpanchyk/ticks2bars/internal/services/handler"
	"github.com/rpanchyk/ticks2bars/internal/services/reader"
	"github.com/rpanchyk/ticks2bars/internal/services/writer"
	"golang.org/x/sync/errgroup"
)

type DefaultPipeline struct {
	config *models.Config
}

func NewPipeline() *DefaultPipeline {
	return &DefaultPipeline{
		config: &globals.Config,
	}
}

func (p *DefaultPipeline) Run(convertable models.Convertable) error {
	fmt.Printf("Convertable: %+v\n", convertable)

	ticksChan := make(chan models.Tick, 1000)
	barsChan := make(chan models.Bar, 1000)

	reader := reader.NewReader(convertable.TicksFile)
	writer := writer.NewWriter(convertable.Symbol, convertable.Timeframes)

	// read
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		defer func() {
			close(ticksChan)
			fmt.Println("Ticks channel closed")
		}()
		return reader.Read(ctx, ticksChan)
	})

	// write
	writer.Init()
	g.Go(func() error {
		defer func() {
			writer.Flush()
			fmt.Println("Writer flushed")
		}()
		return writer.Write(ctx, barsChan)
	})

	// handlers
	var handlersWg sync.WaitGroup
	broker := broker.NewDefaultBroker()
	for _, timeframe := range convertable.Timeframes {
		handler := handler.NewHandler(convertable.Symbol, timeframe, barsChan)
		g.Go(func() error {
			defer func() {
				handler.Flush()
				defer handlersWg.Done()
			}()
			handlersWg.Add(1)
			return handler.Handle(ctx, broker.Subscribe(timeframe))
		})
	}

	// process
	p.readTicks(broker, convertable.Timeframes, ctx, ticksChan)

	// wait for all handlers to finish
	handlersWg.Wait()
	close(barsChan)

	// wait all go routines to finish
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (p *DefaultPipeline) readTicks(broker broker.Broker, timeframes []models.Timeframe, ctx context.Context, ticksChan <-chan models.Tick) error {
	for {
		select {
		case tick, ok := <-ticksChan:
			if !ok {
				for _, timeframe := range timeframes {
					broker.Unsubscribe(timeframe)
				}
				return nil
			}

			broker.Publish(ctx, tick)

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
