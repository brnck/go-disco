package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	LedPin        int    `mapstructure:"led_pin"`
	LedCount      int    `mapstructure:"led_count"`
	LedBrightness int    `mapstructure:"led_brightness"`
	Output        string `mapstructure:"output"`
}

const (
	ledPinDefault        = 18
	ledCountDefault      = 64
	ledBrightnessDefault = 128
	OutputLedStrip       = "ledstrip"
	OutputStdout         = "stdout"
)

func ParseConfig() (*Config, error) {
	var c Config

	setDefaults()
	addConfigTypes()

	if err := addConfigPaths(); err != nil {
		return nil, err
	}
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

func addConfigTypes() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
}

func addConfigPaths() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	viper.AddConfigPath(fmt.Sprintf("%s/config", cwd))
	viper.AddConfigPath("/etc/defaults")
	viper.AddConfigPath("$HOME/src")
	viper.AddConfigPath("$HOME")

	return nil
}

func setDefaults() {
	viper.SetDefault("led_pin", ledPinDefault)
	viper.SetDefault("led_count", ledCountDefault)
	viper.SetDefault("led_brightness", ledBrightnessDefault)
}
