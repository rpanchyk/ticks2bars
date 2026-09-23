package pipeline

import (
	"context"
	"fmt"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
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
			// fmt.Println("ticks channel closed")
		}()
		return reader.Read(ctx, ticksChan)
	})

	// write
	writer.Init()
	g.Go(func() error {
		defer func() {
			writer.Flush()
			// fmt.Println("writer flushed")
		}()
		return writer.Write(ctx, barsChan)
	})

	// process
	p.handle(convertable.Symbol, convertable.Timeframes, ctx, ticksChan, barsChan)

	// wait all go routines to finish
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (p *DefaultPipeline) handle(symbol string, timeframes []models.Timeframe, ctx context.Context, ticksChan <-chan models.Tick, barsChan chan<- models.Bar) error {
	defer func() {
		close(barsChan)
		// fmt.Println("bars channel closed")
	}()

	handlers := make([]handler.Handler, 0, len(timeframes))
	for _, timeframe := range timeframes {
		handlers = append(handlers, handler.NewHandler(symbol, timeframe, barsChan))
	}

	for {
		select {
		case tick, ok := <-ticksChan:
			if !ok {
				for _, handler := range handlers {
					err := handler.Flush()
					if err != nil {
						return err
					}
				}
				return nil
			}
			// fmt.Println("handling tick", tick)

			for _, handler := range handlers {
				err := handler.Handle(tick)
				if err != nil {
					return err
				}
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
