package reader

import (
	"context"

	"github.com/rpanchyk/ticks2bars/internal/models"
)

type Reader interface {
	Read(ctx context.Context, ticksChan chan<- models.Tick) error
}
