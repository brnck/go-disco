package output

import (
	"github.com/brnck/go-disco/app/config"
	"log"
)

type Output interface {
	SetLed(i int, color uint32)
	GetLedColor(i int) uint32
	Render() error
	Open() error
	Close()
}

func InitializeOutput(c *config.Config, l *log.Logger) (Output, error) {
	switch c.Output {
	case config.OutputLedStrip:
		return initializeWS2812Output(c)
	case config.OutputStdout:
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

func initializeWS2812Output(c *config.Config) (*led, error) {
	ws, err := newLed(c)
	if err != nil {
		return nil, err
	}

	return ws, err
}
