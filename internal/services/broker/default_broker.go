package broker

import (
	"context"
	"fmt"
	"sync"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
)

type DefaultBroker struct {
	config      *models.Config
	mu          sync.RWMutex
	subscribers map[models.Timeframe]chan<- models.Tick
}

func NewDefaultBroker() *DefaultBroker {
	return &DefaultBroker{
		config:      &globals.Config,
		subscribers: make(map[models.Timeframe]chan<- models.Tick),
	}
}

func (b *DefaultBroker) Subscribe(timeframe models.Timeframe) chan models.Tick {
	b.mu.Lock()
	defer b.mu.Unlock()

	fmt.Println("Subscribing on timeframe", timeframe)
	ch := make(chan models.Tick, b.config.WriteBatchSize)
	b.subscribers[timeframe] = ch
	return ch
}

func (b *DefaultBroker) Unsubscribe(timeframe models.Timeframe) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.subscribers[timeframe]; exists {
		fmt.Println("Unsubscribing from timeframe", timeframe)
		close(b.subscribers[timeframe])
		delete(b.subscribers, timeframe)
	}
}

func (b *DefaultBroker) Publish(ctx context.Context, tick models.Tick) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, ch := range b.subscribers {
		select {
		case ch <- tick:
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}
