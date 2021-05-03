package output

import (
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
	"log"
)

const (
	strategyLedStrip = "ledstrip"
	strategyStdout   = "stdout"
)

type Output interface {
	SetLed(i int, color uint32)
	GetLedColor(i int) uint32
	Render() error
	Open() error
	Close()
}

func InitializeOutput(name string, c *ws2811.Option, l *log.Logger) (Output, error) {
	switch name {
	case strategyLedStrip:
		return initializeWS2812Output(c)
	case strategyStdout:
		return initializeStandardOutput(l)
	default:
		return nil, errUnsupportedOutput
	}
}

func initializeStandardOutput(l *log.Logger) (*stdout, error) {
	ws, err := newStdout(l)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func initializeWS2812Output(c *ws2811.Option) (*led, error) {
	ws, err := newLed(c)
	if err != nil {
		return nil, err
	}

	return ws, err
}
