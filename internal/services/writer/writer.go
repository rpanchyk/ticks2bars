package writer

import (
	"context"

	"github.com/rpanchyk/ticks2bars/internal/models"
)

type Writer interface {
	Write(ctx context.Context, barsChan <-chan models.Bar) error
}
