package pipeline

import "github.com/rpanchyk/ticks2bars/internal/models"

type Pipeline interface {
	Run(convertable models.Convertable) error
}
