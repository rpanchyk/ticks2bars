package pipeline

import (
	"fmt"

	"github.com/rpanchyk/ticks2bars/internal/models"
	"github.com/rpanchyk/ticks2bars/internal/services/reader"
	"github.com/rpanchyk/ticks2bars/internal/services/writer"
)

type DefaultPipeline struct {
	reader reader.Reader
	writer writer.Writer
}

func NewPipeline(reader reader.Reader, writer writer.Writer) *DefaultPipeline {
	return &DefaultPipeline{
		reader: reader,
		writer: writer,
	}
}

func (p *DefaultPipeline) Run(convertable models.Convertable) error {
	fmt.Printf("Convertable: %+v\n", convertable)

	return nil
}
