package pipeline

import (
	"fmt"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
	"github.com/rpanchyk/ticks2bars/internal/services/reader"
	"github.com/rpanchyk/ticks2bars/internal/services/writer"
)

type DefaultPipeline struct {
	config *models.Config
	reader reader.Reader
	writer writer.Writer
}

func NewPipeline(reader reader.Reader, writer writer.Writer) *DefaultPipeline {
	return &DefaultPipeline{
		config: &globals.Config,
		reader: reader,
		writer: writer,
	}
}

func (p *DefaultPipeline) Run(convertable models.Convertable) error {
	fmt.Printf("Convertable: %+v\n", convertable)

	return nil
}
