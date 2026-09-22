package models

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Tick struct {
	Timestamp time.Time
	Bid       decimal.Decimal
	Ask       decimal.Decimal
}

func (t Tick) String() string {
	return fmt.Sprintf("%s,%s,%s", t.Timestamp, t.Bid, t.Ask)
}
