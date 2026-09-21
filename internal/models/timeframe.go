package models

import (
	"log"
)

type Timeframe string

const (
	m1  Timeframe = "1m"
	m2  Timeframe = "2m"
	m3  Timeframe = "3m"
	m5  Timeframe = "5m"
	m10 Timeframe = "10m"
	m15 Timeframe = "15m"
	m30 Timeframe = "30m"
	h1  Timeframe = "1h"
	h2  Timeframe = "2h"
	h4  Timeframe = "4h"
	h6  Timeframe = "6h"
	h8  Timeframe = "8h"
	h12 Timeframe = "12h"
	D   Timeframe = "D"
	W   Timeframe = "W"
	M   Timeframe = "M"
	Y   Timeframe = "Y"
)

func (t Timeframe) Minutes() int {
	switch t {
	case m1:
		return 1
	case m2:
		return 2
	case m3:
		return 3
	case m5:
		return 5
	case m10:
		return 10
	case m15:
		return 15
	case m30:
		return 30
	case h1:
		return 60
	case h2:
		return 120
	case h4:
		return 240
	case h6:
		return 360
	case h8:
		return 480
	case h12:
		return 720
	case D:
		return 1440
	case W:
		return 10080
	case M:
		return 43200
	case Y:
		return 525600
	default:
		log.Fatal("Unable to get minutes for timeframe:", t)
		return -1
	}
}

func (t Timeframe) String() string {
	return string(t)
}

func CompareTimeframes(a, b Timeframe) int {
	return a.Minutes() - b.Minutes()
}
