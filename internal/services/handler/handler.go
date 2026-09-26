package handler

import (
	"context"

	"github.com/rpanchyk/ticks2bars/internal/models"
)

type Handler interface {
	Handle(ctx context.Context, ticksChan <-chan models.Tick) error
	Flush() error
}
