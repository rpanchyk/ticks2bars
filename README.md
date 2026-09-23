# Ticks to Bars Converter

Utility for converting ticks to bars.

## Configuration

File `config.toml` contains configuration options.

```toml
input_dir = "ticks" # input directory for ticks, --input-dir="ticks"
output_dir = "bars" # output directory for bars, --output-dir="bars"

symbols = "EURUSD,GBPUSD,XAUUSD,NAS100" # symbols, --symbols="EURUSD,GBPUSD,XAUUSD,NAS100"
available_timeframes = "1m,2m,3m,5m,10m,15m,30m,1h,2h,4h,6h,8h,12h,D,W,M,Y" # available timeframes (case-sensitive)
timeframes = "1m,15m,1h" # timeframes (case-sensitive), --timeframes="1m,15m,1h"
include_bars_header = false # include header in bar files, --include-bars-header=false
ticks_timestamp_layout = "2006.01.02 15:04:05.000" # timestamp layout for tick files, --ticks-timestamp-layout="2006.01.02 15:04:05.000"
bars_timestamp_layout = "2006-01-02 15:04:05" # timestamp layout for bar files, --bars-timestamp-layout="2006-01-02 15:04:05"
decimals = 5 # number of decimal places (symbol digits), --decimals=5

force = false # overwrite existing bars, --force=false
write_batch_size = 1000 # batch size for writing data, --write-batch-size=1000
```

## Usage

The following is shown if `ticks2bars --help` executed:

```bash
Usage:
  ticks2bars convert [flags]

Flags:
      --bars-timestamp-layout string    Bars timestamp layout
      --decimals int                    Number of decimal places (symbol digits)
      --force                           Overwrite existing bars
  -h, --help                            help for convert
      --include-bars-header             Include header in bars file
      --input-dir string                Input directory for ticks
      --output-dir string               Output directory for bars
      --symbols string                  Symbols to convert
      --ticks-timestamp-layout string   Ticks timestamp layout
      --timeframes string               Timeframes to convert
      --write-batch-size int            Write batch size

Global Flags:
  -c, --config string   Config file (default "config.toml")
```

For example, the following command:

```bash
ticks2bars convert --symbols="EURUSD" --timeframes="1h"
```

Will convert ticks to bars for the `EURUSD` symbol and `1h` timeframe.

## How it works

Having a `tick` file with the following content:

```csv
Timestamp,Bid,Ask
2020.06.22 00:00:01.397,1.11835,1.11900
2020.06.22 00:00:04.021,1.11841,1.11874
2020.06.22 00:00:22.951,1.11840,1.11873
2020.06.22 00:00:53.614,1.11840,1.11872
2020.06.22 00:00:58.259,1.11840,1.11868
2020.06.22 00:01:01.447,1.11840,1.11871
2020.06.22 00:01:01.498,1.11840,1.11868
2020.06.22 00:01:08.392,1.11840,1.11865
2020.06.22 00:01:25.077,1.11840,1.11862
2020.06.22 00:02:24.336,1.11831,1.11855
2020.06.22 00:02:30.658,1.11819,1.11873
2020.06.22 00:02:42.244,1.11831,1.11884
2020.06.22 00:02:50.782,1.11831,1.11884
2020.06.22 00:02:57.916,1.11831,1.11870
```

Will be converted to a `bar` file (for example, 1 minute bars) with the following content:

```csv
Timestamp,Open,High,Low,Close
2020-06-22 00:00:00,1.11835,1.11841,1.11833,1.11840
2020-06-22 00:01:00,1.11840,1.11840,1.11840,1.11840
2020-06-22 00:02:00,1.11831,1.11831,1.11819,1.11831
```

## Disclaimer

The software is provided "as is", without warranty of any kind, express or
implied, including but not limited to the warranties of merchantability,
fitness for a particular purpose and noninfringement. in no event shall the
authors or copyright holders be liable for any claim, damages or other
liability, whether in an action of contract, tort or otherwise, arising from,
out of or in connection with the software or the use or other dealings in the
software.

## Contribution

If you have any ideas or inspiration for contributing the project,
please create an [issue](https://github.com/rpanchyk/ticks2bars/issues/new)
or a [pull request](https://github.com/rpanchyk/ticks2bars/pulls).
