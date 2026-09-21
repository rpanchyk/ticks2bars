package converter

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/rpanchyk/ticks2bars/internal/models"
	"github.com/rpanchyk/ticks2bars/internal/utils"
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
	fmt.Println("Converter started")

	fmt.Printf("Config: %+v\n", c.config)

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
	var timeframes []string
	for _, s := range strings.Split(c.config.Timeframes, ",") {
		trimmed := strings.TrimSpace(s)
		if slices.Contains(availableTimeframes, trimmed) {
			timeframes = append(timeframes, strings.TrimSpace(s))
		} else {
			fmt.Println("Timeframe is not available:", trimmed)
			os.Exit(1)
		}
	}
	fmt.Printf("Timeframes: %+v\n", timeframes)

	convertables := make(map[string]models.Convertable)
	for _, symbol := range symbols {
		// check if ticks file exists
		ticksFile := filepath.Join(c.config.InputDir, symbol+"_ticks.csv")
		if utils.FileDoesNotExist(ticksFile) {
			fmt.Println("Ticks file not found:", ticksFile)
			os.Exit(1)
		}

		// check if bars file already exists for timeframe
		if !c.config.Force {
			for _, timeframe := range timeframes {
				barsFile := filepath.Join(c.config.OutputDir, symbol+"_"+timeframe+".csv")
				if utils.FileExists(barsFile) {
					fmt.Println(symbol, timeframe, "bars file already exists at", barsFile, "(use --force to overwrite)")
					os.Exit(1)
				}
			}
		}

		convertables[symbol] = models.Convertable{
			TicksFile:  ticksFile,
			Timeframes: timeframes,
		}
	}
	for k, v := range convertables {
		fmt.Printf("Convertable %s: %+v\n", k, v)
	}

	fmt.Println("Converter finished")
	return nil
}
