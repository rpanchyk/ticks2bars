package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config string
)

var rootCmd = &cobra.Command{
	Use:   "ticks2bars",
	Short: "Ticks to Bars Converter",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "config.toml", "Config file")
}

func initConfig() {
	configFilePath := config
	if !filepath.IsAbs(config) {
		// executable path
		exePath, err := os.Executable()
		if err != nil {
			fmt.Println("Cannot get executable path, error:", err)
			os.Exit(1)
		}
		configFilePath = filepath.Join(filepath.Dir(exePath), config)

		// current directory
		if utils.FileDoesNotExist(configFilePath) {
			fmt.Println("Config not found at", configFilePath)
			currDir, err := os.Getwd()
			if err != nil {
				fmt.Println("Cannot get current directory, error:", err)
				os.Exit(1)
			}
			configFilePath = filepath.Join(currDir, config)
		}

		// user home directory
		if utils.FileDoesNotExist(configFilePath) {
			fmt.Println("Config not found at", configFilePath)
			userHomeDir, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Cannot get user home directory, error:", err)
				os.Exit(1)
			}
			configFilePath = filepath.Join(userHomeDir, config)
		}
	}

	if utils.FileDoesNotExist(configFilePath) {
		fmt.Println("Config not found at", configFilePath)
		os.Exit(1)
	}
	fmt.Println("Config found at", configFilePath)

	configDir := filepath.Dir(configFilePath)
	configFile := filepath.Base(configFilePath)
	configExt := filepath.Ext(configFile)
	configName := strings.TrimSuffix(configFile, configExt)

	viper.AddConfigPath(configDir)
	viper.SetConfigName(configName)
	viper.SetConfigType(configExt[1:])
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Cannot read config, error:", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&globals.Config); err != nil {
		fmt.Println("Cannot unmarshal config, error:", err)
		os.Exit(1)
	}
	fmt.Printf("Config: %+v\n", globals.Config)
}
