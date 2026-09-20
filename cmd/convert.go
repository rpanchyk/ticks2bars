package cmd

import (
	"fmt"
	"os"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
	"github.com/rpanchyk/ticks2bars/internal/services/converter"
	"github.com/rpanchyk/ticks2bars/internal/utils"
	"github.com/spf13/cobra"
)

var (
	inputDir             string
	outputDir            string
	symbols              string
	timeframes           string
	includeHeader        bool
	ticksTimestampLayout string
	barsTimestampLayout  string
	force                bool
	writeBatchSize       int
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert ticks to bars",
	Run: func(cmd *cobra.Command, args []string) {

		config := models.Config{
			InputDir:             utils.ToAbsPath(inputDir),
			OutputDir:            utils.ToAbsPath(outputDir),
			Symbols:              symbols,
			AvailableTimeframes:  globals.Config.AvailableTimeframes,
			Timeframes:           timeframes,
			IncludeHeader:        includeHeader,
			TicksTimestampLayout: ticksTimestampLayout,
			BarsTimestampLayout:  barsTimestampLayout,
			Force:                force,
			WriteBatchSize:       writeBatchSize,
		}

		converter := converter.NewConverter(&config)

		err := converter.Convert()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	convertCmd.Flags().StringVar(&inputDir, "input-dir", globals.Config.InputDir, "Input directory for ticks")
	convertCmd.Flags().StringVar(&outputDir, "output-dir", globals.Config.OutputDir, "Output directory for bars")
	convertCmd.Flags().StringVar(&symbols, "symbols", globals.Config.Symbols, "Symbols to convert")
	convertCmd.Flags().StringVar(&timeframes, "timeframes", globals.Config.Timeframes, "Timeframes to convert")
	convertCmd.Flags().BoolVar(&includeHeader, "include-header", globals.Config.IncludeHeader, "Include header in output")
	convertCmd.Flags().StringVar(&ticksTimestampLayout, "ticks-timestamp-layout", globals.Config.TicksTimestampLayout, "Ticks timestamp layout")
	convertCmd.Flags().StringVar(&barsTimestampLayout, "bars-timestamp-layout", globals.Config.BarsTimestampLayout, "Bars timestamp layout")
	convertCmd.Flags().BoolVar(&force, "force", globals.Config.Force, "Overwrite existing bars")
	convertCmd.Flags().IntVar(&writeBatchSize, "write-batch-size", globals.Config.WriteBatchSize, "Write batch size")

	rootCmd.AddCommand(convertCmd)
}
