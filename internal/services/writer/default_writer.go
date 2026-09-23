package writer

import (
	"context"
	"os"
	"path/filepath"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
)

type DefaultWriter struct {
	config     *models.Config
	symbol     string
	timeframes []models.Timeframe
	bars       map[models.Timeframe][]models.Bar
}

func NewWriter(symbol string, timeframes []models.Timeframe) *DefaultWriter {
	return &DefaultWriter{
		config:     &globals.Config,
		symbol:     symbol,
		timeframes: timeframes,
		bars:       nil,
	}
}

func (w *DefaultWriter) Init() error {
	// initialize bars
	w.bars = make(map[models.Timeframe][]models.Bar, len(w.timeframes))
	// bars := map[models.Timeframe][]models.Bar{}
	// for _, timeframe := range w.timeframes {
	// 	bars[timeframe] = make([]models.Bar, 0, w.config.WriteBatchSize)
	// }
	// w.bars = bars

	// initialize files
	for _, timeframe := range w.timeframes {
		err := w.createBarsFile(w.symbol, timeframe)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *DefaultWriter) Write(ctx context.Context, barsChan <-chan models.Bar) error {
	for {
		select {
		case bar, ok := <-barsChan:
			if !ok {
				return nil
			}
			// fmt.Println("writing bar", bar)

			tf := bar.Timeframe
			if _, exists := w.bars[tf]; !exists {
				w.bars[tf] = make([]models.Bar, 0, w.config.WriteBatchSize)
			}

			if len(w.bars[tf]) < w.config.WriteBatchSize {
				w.bars[tf] = append(w.bars[tf], bar)
			} else {
				err := w.writeBarsFile(bar.Symbol, bar.Timeframe, w.bars[tf])
				if err != nil {
					return err
				}
				w.bars[tf] = make([]models.Bar, 0, w.config.WriteBatchSize)
				w.bars[tf] = append(w.bars[bar.Timeframe], bar)
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (w *DefaultWriter) Flush() error {
	for tf, bars := range w.bars {
		err := w.writeBarsFile(w.symbol, tf, bars)
		if err != nil {
			return err
		}
		w.bars[tf] = make([]models.Bar, 0, w.config.WriteBatchSize)
	}
	return nil
}

func (w *DefaultWriter) createBarsFile(symbol string, timeframe models.Timeframe) error {
	filepath := w.getBarsFilePath(symbol, timeframe)

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	content := ""
	if w.config.IncludeBarsHeader {
		content += "Timestamp,Open,High,Low,Close\n"
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func (w *DefaultWriter) writeBarsFile(symbol string, timeframe models.Timeframe, bars []models.Bar) error {
	filepath := w.getBarsFilePath(symbol, timeframe)

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	content := ""
	for _, bar := range bars {
		content += bar.Timestamp.Format(w.config.BarsTimestampLayout)
		content += "," + bar.Open.StringFixed(5)
		content += "," + bar.High.StringFixed(5)
		content += "," + bar.Low.StringFixed(5)
		content += "," + bar.Close.StringFixed(5)
		content += "\n"
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func (w *DefaultWriter) getBarsFilePath(symbol string, timeframe models.Timeframe) string {
	return filepath.Join(w.config.OutputDir, symbol+"_"+timeframe.String()+"_"+".csv")
}
