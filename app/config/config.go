// Package config is implemented to path configuration to the application.
package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
)

// An App is a global application configuration instance.
var (
	App Config
)

// Represents app configuration type.
type Config struct {
	Board    BoardConfig
	LogLevel string `env:"LOG_LEVEL" env-default:"debug"`
}

// Represents game board configuration type.
type BoardConfig struct {
	Width           int `env:"BOARD_WIDTH" env-default:"10"`
	WidthMax        int `env:"BOARD_WIDTH_MAX" env-default:"10"`
	Height          int `env:"BOARD_HEIGHT" env-default:"10"`
	HeightMax       int `env:"BOARD_HEIGHT_MAX" env-default:"10"`
	BlackHolesCount int `env:"BOARD_BLACK_HOLES_COUNT" env-default:"5"`
}

// Loads configuration from the config file into environment variables.
func (c *Config) Load() {
	if err := env.Parse(&App); err != nil {
		fmt.Printf("Could not load config: %v\n", err)
		os.Exit(-1)
	}
}
