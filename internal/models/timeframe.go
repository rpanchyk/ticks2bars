package models

import (
	"log"
)

type Timeframe string

const (
	TF_m1  Timeframe = "1m"
	TF_m2  Timeframe = "2m"
	TF_m3  Timeframe = "3m"
	TF_m5  Timeframe = "5m"
	TF_m10 Timeframe = "10m"
	TF_m15 Timeframe = "15m"
	TF_m30 Timeframe = "30m"
	TF_h1  Timeframe = "1h"
	TF_h2  Timeframe = "2h"
	TF_h4  Timeframe = "4h"
	TF_h6  Timeframe = "6h"
	TF_h8  Timeframe = "8h"
	TF_h12 Timeframe = "12h"
	TF_D   Timeframe = "D"
	TF_W   Timeframe = "W"
	TF_M   Timeframe = "M"
	TF_Y   Timeframe = "Y"
)

func (t Timeframe) Minutes() int {
	switch t {
	case TF_m1:
		return 1
	case TF_m2:
		return 2
	case TF_m3:
		return 3
	case TF_m5:
		return 5
	case TF_m10:
		return 10
	case TF_m15:
		return 15
	case TF_m30:
		return 30
	case TF_h1:
		return 60
	case TF_h2:
		return 120
	case TF_h4:
		return 240
	case TF_h6:
		return 360
	case TF_h8:
		return 480
	case TF_h12:
		return 720
	case TF_D:
		return 1440
	case TF_W:
		return 10080
	case TF_M:
		return 43200 // approximate
	case TF_Y:
		return 525600 // approximate
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
