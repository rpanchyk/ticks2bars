package converter

import (
	"fmt"

	"github.com/rpanchyk/ticks2bars/internal/models"
)

type DefaultConverter struct {
	config *models.Config
}

func NewConverter(config *models.Config) *DefaultConverter {
	return &DefaultConverter{
		config: config,
	}
}

func (c *DefaultConverter) Convert() error {
	fmt.Println("Default converter started")

	fmt.Printf("Config: %+v\n", c.config)

	//...

	fmt.Println("Default converter finished")
	return nil
}
