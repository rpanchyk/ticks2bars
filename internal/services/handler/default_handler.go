package handler

import (
	"fmt"
	"time"

	"github.com/rpanchyk/ticks2bars/internal/models"
	"github.com/shopspring/decimal"
)

type DefaultHandler struct {
	barsChan chan<- models.Bar
	bar      models.Bar
}

func NewHandler(symbol string, timeframe models.Timeframe, barsChan chan<- models.Bar) *DefaultHandler {
	return &DefaultHandler{
		barsChan: barsChan,
		bar: models.Bar{
			Symbol:    symbol,
			Timeframe: timeframe,
		},
	}
}

func (h *DefaultHandler) Handle(tick models.Tick) error {
	// fmt.Println(h.bar.Timeframe, "time=", tick.Timestamp)
	interval := time.Duration(h.bar.Timeframe.Minutes()) * time.Minute
	barTime := tick.Timestamp.Truncate(interval)

	// adjust monthly and yearly bars
	switch h.bar.Timeframe {
	case models.TF_M:
		barTime = time.Date(tick.Timestamp.Year(), tick.Timestamp.Month(), 1, 0, 0, 0, 0, tick.Timestamp.Location())
	case models.TF_Y:
		barTime = time.Date(tick.Timestamp.Year(), 1, 1, 0, 0, 0, 0, tick.Timestamp.Location())
	}
	// fmt.Println(h.bar.Timeframe, "barTime=", barTime)

	// price := (tick.Bid.Add(tick.Ask)).Div(decimal.NewFromInt(2))
	price := tick.Bid

	if isBarEmpty(h.bar) || isBarComplete(h.bar, barTime) {
		err := h.Flush()
		if err != nil {
			return err
		}

		// new bar
		h.bar.Timestamp = barTime
		h.bar.Open = price
		h.bar.High = price
		h.bar.Low = price
		h.bar.Close = price
	} else {
		// update bar
		h.bar.High = decimal.Max(h.bar.High, price)
		h.bar.Low = decimal.Min(h.bar.Low, price)
		h.bar.Close = price
	}
	return nil
}

func (h *DefaultHandler) Flush() error {
	if !isBarEmpty(h.bar) {
		fmt.Println("sending bar", h.bar)
		h.barsChan <- h.bar
	}
	return nil
}

func isBarEmpty(bar models.Bar) bool {
	return bar.Timestamp.Equal(time.Time{})
}

func isBarComplete(bar models.Bar, barTime time.Time) bool {
	return bar.Timestamp.Before(barTime)
}
