package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Tick struct {
	Timestamp time.Time
	Bid       decimal.Decimal
	Ask       decimal.Decimal
}
