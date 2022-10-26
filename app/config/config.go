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
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

// Represents game board configuration type.
type BoardConfig struct {
	Width           int `env:"BOARD_WIDTH" envDefault:"8"`
	WidthMin        int `env:"BOARD_WIDTH_MIN" envDefault:"3"`
	WidthMax        int `env:"BOARD_WIDTH_MAX" envDefault:"10"`
	Height          int `env:"BOARD_HEIGHT" envDefault:"8"`
	HeightMin       int `env:"BOARD_HEIGHT_MIN" envDefault:"3"`
	HeightMax       int `env:"BOARD_HEIGHT_MAX" envDefault:"10"`
	BlackHolesCount int `env:"BOARD_BLACK_HOLES_COUNT" envDefault:"10"`
}

// Loads configuration from the config file into environment variables.
func (c *Config) Load() {
	if err := env.Parse(&App); err != nil {
		fmt.Printf("Could not load config: %v\n", err)
		os.Exit(-1)
	}
}
