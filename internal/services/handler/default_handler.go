package handler

import (
	"fmt"

	"github.com/rpanchyk/ticks2bars/internal/models"
)

type DefaultHandler struct {
	symbol    string
	timeframe models.Timeframe
	barsChan  chan<- models.Bar
}

func NewHandler(symbol string, timeframe models.Timeframe, barsChan chan<- models.Bar) *DefaultHandler {
	return &DefaultHandler{
		symbol:    symbol,
		timeframe: timeframe,
		barsChan:  barsChan,
	}
}

func (h *DefaultHandler) Handle(tick models.Tick) error {
	bar := models.Bar{
		Symbol:    h.symbol,
		Timeframe: h.timeframe,
		Timestamp: "2022-01-01 00:00:00",
		Open:      tick.Bid.String(),
		High:      tick.Bid.String(),
		Low:       tick.Ask.String(),
		Close:     tick.Ask.String(),
	}

	fmt.Println("sending bar", bar)
	h.barsChan <- bar

	return nil
}
