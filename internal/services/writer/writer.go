package writer

import "github.com/rpanchyk/ticks2bars/internal/models"

type Writer interface {
	Write(bars []models.Bar) error
}
