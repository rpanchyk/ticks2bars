package utils

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func ToAbsPath(path string) string {
	if !filepath.IsAbs(path) {
		configFile := viper.GetViper().ConfigFileUsed()
		return filepath.Join(filepath.Dir(configFile), path)
	}
	return path
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func FileDoesNotExist(filePath string) bool {
	return !FileExists(filePath)
}
