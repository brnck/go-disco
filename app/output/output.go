package output

import (
	"log"
)

type Output interface {
	SetLed(i int, color uint32)
	Render() error
	Open() error
	Close()
}

func InitializeStandardOutput(l *log.Logger) (*stdout, error) {
	ws, err := NewStdout(l)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func InitializeWS2812Output() (*led, error) {
	ws, err := NewLed()
	if err != nil {
		return nil, err
	}

	return ws, err
}
