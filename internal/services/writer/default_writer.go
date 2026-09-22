package writer

import (
	"context"
	"fmt"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
)

type DefaultWriter struct {
	config *models.Config
}

func NewWriter() *DefaultWriter {
	return &DefaultWriter{
		config: &globals.Config,
	}
}

func (w *DefaultWriter) Write(ctx context.Context, barsChan <-chan models.Bar) error {
	for {
		select {
		case bar, ok := <-barsChan:
			if !ok {
				return nil
			}
			fmt.Println("writing bar", bar)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
