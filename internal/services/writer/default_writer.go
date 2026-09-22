package writer

import (
	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
)

type DefaultWriter struct {
	config *models.Config
}

func NewWriter() *DefaultWriter {
	return &DefaultWriter{
		config: &globals.Config,
	}
}

func (w *DefaultWriter) Write(bars []models.Bar) error {
	return w.Write(bars)
}
