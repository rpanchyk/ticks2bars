package pipeline

import (
	"context"
	"fmt"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
	"github.com/rpanchyk/ticks2bars/internal/services/reader"
	"github.com/rpanchyk/ticks2bars/internal/services/writer"
	"golang.org/x/sync/errgroup"
)

type DefaultPipeline struct {
	config *models.Config
	reader reader.Reader
	writer writer.Writer
}

func NewPipeline(reader reader.Reader, writer writer.Writer) *DefaultPipeline {
	return &DefaultPipeline{
		config: &globals.Config,
		reader: reader,
		writer: writer,
	}
}

func (p *DefaultPipeline) Run(convertable models.Convertable) error {
	fmt.Printf("Convertable: %+v\n", convertable)

	ticksChan := make(chan models.Tick, 1000)
	barsChan := make(chan models.Bar, 1000)

	// read
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		defer close(ticksChan)
		return p.reader.Read(convertable.TicksFile, ctx, ticksChan)
	})

	// handle
	p.handle(ctx, ticksChan, barsChan)

	// write
	g.Go(func() error {
		defer close(barsChan)
		return p.writer.Write(ctx, barsChan)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Done")
	}
	return nil
}

func (p *DefaultPipeline) handle(ctx context.Context, ticksChan <-chan models.Tick, barsChan chan<- models.Bar) error {
	for {
		select {
		case tick, ok := <-ticksChan:
			if !ok {
				return nil
			}
			fmt.Println("handling tick", tick)

			bar := models.Bar{
				Symbol:    "BTCUSDT",
				Timeframe: models.Timeframe("1m"),
				Timestamp: tick.Timestamp.String(),
				Open:      tick.Bid.String(),
				High:      tick.Bid.String(),
				Low:       tick.Ask.String(),
				Close:     tick.Ask.String(),
			}
			barsChan <- bar
			fmt.Println("sending bar", bar)

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
