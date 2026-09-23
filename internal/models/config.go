package models

type Config struct {
	InputDir  string `mapstructure:"input_dir"`
	OutputDir string `mapstructure:"output_dir"`

	Symbols             string `mapstructure:"symbols"`
	AvailableTimeframes string `mapstructure:"available_timeframes"`
	Timeframes          string `mapstructure:"timeframes"`

	IncludeBarsHeader    bool   `mapstructure:"include_bars_header"`
	TicksTimestampLayout string `mapstructure:"ticks_timestamp_layout"`
	BarsTimestampLayout  string `mapstructure:"bars_timestamp_layout"`
	Decimals             int    `mapstructure:"decimals"`

	Force          bool `mapstructure:"force"`
	WriteBatchSize int  `mapstructure:"write_batch_size"`
}
