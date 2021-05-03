package app

import (
	"fmt"
	"github.com/brnck/go-disco/pkg/program"
	"github.com/brnck/go-disco/pkg/scene"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	LedPin        int            `mapstructure:"led_pin"`
	LedCount      int            `mapstructure:"led_count"`
	LedBrightness int            `mapstructure:"led_brightness"`
	Output        string         `mapstructure:"output"`
	Scenes        []scene.Config `mapstructure:"scenes"`
}

const (
	ledPinDefault        = 18
	ledCountDefault      = 64
	ledBrightnessDefault = 128
)

func parseConfig() (*Config, error) {
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

	viper.SetDefault("iterations", program.IterationsDefault)
	viper.SetDefault("wait_time", program.WaitTimeDefault)
	viper.SetDefault("speed", program.SpeedDefault)
	viper.SetDefault("red", program.RedDefault)
	viper.SetDefault("green", program.GreenDefault)
	viper.SetDefault("blue", program.BlueDefault)
	viper.SetDefault("map", make([]program.Map, 0))

	viper.SetDefault("random_decay", program.RandomDecayDefault)
	viper.SetDefault("trail_decay", program.TrailDecayDefault)
	viper.SetDefault("size", program.SizeDefault)

	viper.SetDefault("one_led_per_scene", program.OneLedPerSceneDefault)

	viper.SetDefault("heat_key", program.HeatKeyDefault)
	viper.SetDefault("cooling", program.CoolingDefault)
	viper.SetDefault("sparking", program.SparkingDefault)
}

func buildOutputConfig(c *Config) *ws2811.Option {
	options := &ws2811.DefaultOptions
	options.Channels[0].GpioPin = c.LedPin
	options.Channels[0].LedCount = c.LedCount
	options.Channels[0].Brightness = c.LedBrightness

	return options
}
