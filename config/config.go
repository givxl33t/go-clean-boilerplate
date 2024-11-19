package config

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

func New() *viper.Viper {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	configFile := path.Join(currentDir, "..", ".env")

	viper := viper.New()
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	// Try loading the .env file, but don't panic if it's missing for testing
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Warning: .env file not found or could not be loaded: %v\n", err)
	}

	return viper
}
