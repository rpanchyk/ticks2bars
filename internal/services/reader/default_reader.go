package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
	"github.com/shopspring/decimal"
)

type DefaultReader struct {
	config *models.Config
}

func NewReader() *DefaultReader {
	return &DefaultReader{
		config: &globals.Config,
	}
}

func (r *DefaultReader) Read(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Cannot open file:", err)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lineNumber := 0
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading line %d: %v\n", lineNumber, err)
			return err
		}

		timestamp, err := time.Parse(r.config.TicksTimestampLayout, record[0])
		if err != nil {
			fmt.Printf("Error parsing timestamp: %v\n", err)
			return err
		}
		bid, err := decimal.NewFromString(record[1])
		if err != nil {
			return err
		}
		ask, err := decimal.NewFromString(record[2])
		if err != nil {
			return err
		}

		tick := models.Tick{Timestamp: timestamp, Bid: bid, Ask: ask}
		fmt.Printf("Tick %d: %v\n", lineNumber, tick)

		if lineNumber > 5 {
			break
		}
		lineNumber++
	}

	return nil
}
