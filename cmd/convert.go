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
			InputDir:             utils.ToAbsPath(utils.FirstNonEmptyString(inputDir, globals.Config.InputDir)),
			OutputDir:            utils.ToAbsPath(utils.FirstNonEmptyString(outputDir, globals.Config.OutputDir)),
			Symbols:              utils.FirstNonEmptyString(symbols, globals.Config.Symbols),
			AvailableTimeframes:  globals.Config.AvailableTimeframes,
			Timeframes:           utils.FirstNonEmptyString(timeframes, globals.Config.Timeframes),
			IncludeHeader:        utils.FirstNonFalseBool(includeHeader, globals.Config.IncludeHeader),
			TicksTimestampLayout: utils.FirstNonEmptyString(ticksTimestampLayout, globals.Config.TicksTimestampLayout),
			BarsTimestampLayout:  utils.FirstNonEmptyString(barsTimestampLayout, globals.Config.BarsTimestampLayout),
			Force:                utils.FirstNonFalseBool(force, globals.Config.Force),
			WriteBatchSize:       utils.FirstNonZeroInt(writeBatchSize, globals.Config.WriteBatchSize),
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
	convertCmd.Flags().StringVar(&inputDir, "input-dir", "", "Input directory for ticks")
	convertCmd.Flags().StringVar(&outputDir, "output-dir", "", "Output directory for bars")
	convertCmd.Flags().StringVar(&symbols, "symbols", "", "Symbols to convert")
	convertCmd.Flags().StringVar(&timeframes, "timeframes", "", "Timeframes to convert")
	convertCmd.Flags().BoolVar(&includeHeader, "include-header", false, "Include header in output")
	convertCmd.Flags().StringVar(&ticksTimestampLayout, "ticks-timestamp-layout", "", "Ticks timestamp layout")
	convertCmd.Flags().StringVar(&barsTimestampLayout, "bars-timestamp-layout", "", "Bars timestamp layout")
	convertCmd.Flags().BoolVar(&force, "force", false, "Overwrite existing bars")
	convertCmd.Flags().IntVar(&writeBatchSize, "write-batch-size", 0, "Write batch size")

	rootCmd.AddCommand(convertCmd)
}
