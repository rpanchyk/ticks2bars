package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
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
		os.Exit(1)
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
			break
		}

		lineNumber++
		// fmt.Printf("Line %d: %v\n", lineNumber, record)

		// dateStr := "25.12.2026 18:30:00"

		// // Створюємо кастомний layout, використовуючи опорну дату Go
		// // Опорний шаблон для "ДД.ММ.РРРР ГГ:ХХ:СС" виглядатиме так:
		// layout := "02.01.2006 15:04:05"

		// parsedTime, err := time.Parse(layout, dateStr)
		// if err != nil {
		// 	fmt.Printf("Помилка парсингу дати: %v\n", err)
		// 	return
		// }

		tick := models.Tick{Timestamp: record[0], Bid: record[1], Ask: record[2]}
		fmt.Printf("Tick \t\t%d: %v\n", lineNumber, tick)

		if lineNumber > 5 {
			break
		}
	}

	return nil
}
