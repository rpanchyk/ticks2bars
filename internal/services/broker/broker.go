package broker

import (
	"context"

	"github.com/rpanchyk/ticks2bars/internal/models"
)

type Broker interface {
	Subscribe(timeframe models.Timeframe) chan models.Tick
	Unsubscribe(timeframe models.Timeframe)
	Publish(ctx context.Context, tick models.Tick) error
}
