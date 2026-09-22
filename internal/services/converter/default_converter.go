package converter

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
	"github.com/rpanchyk/ticks2bars/internal/services/pipeline"
	"github.com/rpanchyk/ticks2bars/internal/utils"
)

type DefaultConverter struct {
	config   *models.Config
	pipeline pipeline.Pipeline
}

func NewConverter(pipeline pipeline.Pipeline) *DefaultConverter {
	return &DefaultConverter{
		config:   &globals.Config,
		pipeline: pipeline,
	}
}

func (c *DefaultConverter) Convert() error {
	fmt.Println("Converter started")

	// symbols
	var symbols []string
	for _, s := range strings.Split(c.config.Symbols, ",") {
		symbols = append(symbols, strings.ToUpper(strings.TrimSpace(s)))
	}
	fmt.Printf("Symbols: %+v\n", symbols)

	// available timeframes
	var availableTimeframes []string
	for _, s := range strings.Split(c.config.AvailableTimeframes, ",") {
		availableTimeframes = append(availableTimeframes, strings.TrimSpace(s))
	}
	fmt.Printf("Available timeframes: %+v\n", availableTimeframes)

	// timeframes
	var timeframes []models.Timeframe
	for _, timeframe := range strings.Split(c.config.Timeframes, ",") {
		trimmed := strings.TrimSpace(timeframe)
		if slices.Contains(availableTimeframes, trimmed) {
			timeframes = append(timeframes, models.Timeframe(trimmed))
		} else {
			return fmt.Errorf("Timeframe is not available: %s", trimmed)
		}
	}
	slices.SortFunc(timeframes, models.CompareTimeframes)
	fmt.Printf("Timeframes: %+v\n", timeframes)

	// prepare
	convertables := []models.Convertable{}
	for _, symbol := range symbols {
		// check if ticks file exists
		ticksFile := filepath.Join(c.config.InputDir, symbol+"_ticks.csv")
		if utils.FileDoesNotExist(ticksFile) {
			return fmt.Errorf("Ticks file not found: %s", ticksFile)
		}

		// check if bars file already exists for timeframe
		for _, timeframe := range timeframes {
			barsFile := filepath.Join(c.config.OutputDir, symbol+"_"+timeframe.String()+".csv")
			if utils.FileExists(barsFile) {
				if !c.config.Force {
					return fmt.Errorf(symbol, timeframe, "bars file already exists at", barsFile, "(use --force to overwrite)")
				} else {
					fmt.Println(symbol, timeframe, "bars file already exists at", barsFile, "(will be overwritten)")
				}
			}
		}

		// add
		convertables = append(convertables, models.Convertable{
			TicksFile:  ticksFile,
			Symbol:     symbol,
			Timeframes: timeframes,
		})
	}

	// run
	for _, convertable := range convertables {
		err := c.pipeline.Run(convertable)
		if err != nil {
			return err
		}
	}

	fmt.Println("Converter finished")
	return nil
}
