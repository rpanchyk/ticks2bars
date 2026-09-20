package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
		fmt.Println("Converting ticks to bars...")
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

	fmt.Println("Trying to get config file...")
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Cannot get executable path, error:", err)
		os.Exit(1)
	}
	configDir := filepath.Dir(exePath)
	fmt.Println("Executable directory:", configDir)

	if isConfigNotFound(configDir) {
		fmt.Println("Config not found, trying to get from current directory...")
		currDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Cannot get current directory, error:", err)
			os.Exit(1)
		}
		configDir = currDir
		fmt.Println("Current directory:", configDir)
	}

	if isConfigNotFound(configDir) {
		fmt.Println("Config not found, trying to get from user home directory...")
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Cannot get user home directory, error:", err)
			os.Exit(1)
		}
		configDir = userHomeDir
		fmt.Println("User home directory:", configDir)
	}

	if isConfigNotFound(configDir)  {
		fmt.Println("Config not found, exiting...")
		os.Exit(1)
	}

	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Cannot read config, error:", err)
		os.Exit(1)
	}

	globals.Config = getConfig()
	fmt.Printf("Config: %+v\n", globals.Config)
}

func isConfigNotFound(exeDir string) bool {
	_, err := os.Stat(filepath.Join(exeDir, globals.ConfigFile))
	if err != nil {
		if os.IsNotExist(err) {
			return true
		} else {
			fmt.Println("Cannot get config file, error:", err)
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
