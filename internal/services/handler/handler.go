package handler

import "github.com/rpanchyk/ticks2bars/internal/models"

type Handler interface {
	Handle(tick models.Tick) error
}
