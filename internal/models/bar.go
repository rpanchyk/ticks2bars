package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Bar struct {
	Symbol    string
	Timeframe Timeframe
	Timestamp time.Time
	Open      decimal.Decimal
	High      decimal.Decimal
	Low       decimal.Decimal
	Close     decimal.Decimal
}
