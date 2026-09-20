package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rpanchyk/ticks2bars/internal/globals"
	"github.com/rpanchyk/ticks2bars/internal/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ticks2bars",
	Short: "Ticks to Bars Converter",
	Long: `Ticks to Bars Converter is a tool that converts ticks to bars.`,
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
}

func initConfig() {
	globals.ConfigFile = "config.toml"

	fmt.Println("Getting config...")
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Cannot get executable path, error:", err)
		os.Exit(1)
	}
	configFilePath := filepath.Join(filepath.Dir(exePath), globals.ConfigFile)

	if isConfigNotFound(configFilePath) {
		fmt.Println("Config not found:", configFilePath)
		currDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Cannot get current directory, error:", err)
			os.Exit(1)
		}
		configFilePath = filepath.Join(currDir, globals.ConfigFile)
	}

	if isConfigNotFound(configFilePath) {
		fmt.Println("Config not found:", configFilePath)
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Cannot get user home directory, error:", err)
			os.Exit(1)
		}
		configFilePath = filepath.Join(userHomeDir, globals.ConfigFile)
	}

	if isConfigNotFound(configFilePath) {
		fmt.Println("Config not found:", configFilePath)
		fmt.Println("Exiting...")
		os.Exit(1)
	}
	fmt.Println("Config found:", configFilePath)

	configPath := filepath.Dir(configFilePath)
	configFile := filepath.Base(configFilePath)
	configExt := filepath.Ext(configFile)
	configName := strings.TrimSuffix(configFile, configExt)

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configExt[1:])
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Cannot read config, error:", err)
		os.Exit(1)
	}

	globals.Config = getConfig()
	fmt.Printf("Config: %+v\n\n", globals.Config)
}

func isConfigNotFound(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return true
		} else {
			fmt.Println("Cannot check file, error:", err)
			os.Exit(1)
		}
	}
	return false
}

func getConfig() models.Config {
	var config models.Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Cannot unmarshal config, error:", err)
		os.Exit(1)
	}
	return config
}
