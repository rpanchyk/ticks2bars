package writer

import "github.com/rpanchyk/ticks2bars/internal/models"

type DefaultWriter struct {
}

func NewWriter() *DefaultWriter {
	return &DefaultWriter{}
}

func (w *DefaultWriter) Write(bars []models.Bar) error {
	return w.Write(bars)
}
