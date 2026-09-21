package cmd

import (
	"fmt"
	"os"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
	"github.com/rpanchyk/ticks2bars/internal/services/converter"
	"github.com/rpanchyk/ticks2bars/internal/services/pipeline"
	"github.com/rpanchyk/ticks2bars/internal/services/reader"
	"github.com/rpanchyk/ticks2bars/internal/services/writer"
	"github.com/rpanchyk/ticks2bars/internal/utils"
	"github.com/spf13/cobra"
)

var (
	inputDir             string
	outputDir            string
	symbols              string
	timeframes           string
	includeBarsHeader    bool
	ticksTimestampLayout string
	barsTimestampLayout  string
	force                bool
	writeBatchSize       int
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert ticks to bars",
	Run: func(cmd *cobra.Command, args []string) {
		config := &models.Config{
			InputDir:             utils.ToAbsPath(utils.FirstNonEmptyString(inputDir, globals.Config.InputDir)),
			OutputDir:            utils.ToAbsPath(utils.FirstNonEmptyString(outputDir, globals.Config.OutputDir)),
			Symbols:              utils.FirstNonEmptyString(symbols, globals.Config.Symbols),
			AvailableTimeframes:  globals.Config.AvailableTimeframes,
			Timeframes:           utils.FirstNonEmptyString(timeframes, globals.Config.Timeframes),
			IncludeBarsHeader:    utils.FirstNonFalseBool(includeBarsHeader, globals.Config.IncludeBarsHeader),
			TicksTimestampLayout: utils.FirstNonEmptyString(ticksTimestampLayout, globals.Config.TicksTimestampLayout),
			BarsTimestampLayout:  utils.FirstNonEmptyString(barsTimestampLayout, globals.Config.BarsTimestampLayout),
			Force:                utils.FirstNonFalseBool(force, globals.Config.Force),
			WriteBatchSize:       utils.FirstNonZeroInt(writeBatchSize, globals.Config.WriteBatchSize),
		}
		reader := reader.NewReader()
		writer := writer.NewWriter()
		pipeline := pipeline.NewPipeline(reader, writer)
		converter := converter.NewConverter(config, pipeline)

		err := converter.Convert()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	convertCmd.Flags().StringVar(&inputDir, "input-dir", "", "Input directory for ticks")
	convertCmd.Flags().StringVar(&outputDir, "output-dir", "", "Output directory for bars")
	convertCmd.Flags().StringVar(&symbols, "symbols", "", "Symbols to convert")
	convertCmd.Flags().StringVar(&timeframes, "timeframes", "", "Timeframes to convert")
	convertCmd.Flags().BoolVar(&includeBarsHeader, "include-bars-header", false, "Include header in bars file")
	convertCmd.Flags().StringVar(&ticksTimestampLayout, "ticks-timestamp-layout", "", "Ticks timestamp layout")
	convertCmd.Flags().StringVar(&barsTimestampLayout, "bars-timestamp-layout", "", "Bars timestamp layout")
	convertCmd.Flags().BoolVar(&force, "force", false, "Overwrite existing bars")
	convertCmd.Flags().IntVar(&writeBatchSize, "write-batch-size", 0, "Write batch size")

	rootCmd.AddCommand(convertCmd)
}
