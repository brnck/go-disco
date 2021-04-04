package output

import (
	"github.com/brnck/go-disco/app/config"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

type led struct {
	ws *ws2811.WS2811
}

func (l *led) SetLed(i int, color uint32) {
	l.ws.Leds(0)[i] = color
}

func (l *led) Render() error {
	return l.ws.Render()
}

func (l *led) Open() error {
	return l.ws.Init()
}

func (l *led) Close() {
	l.ws.Fini()
}

func newLed(c *config.Config) (*led, error) {
	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = c.LedBrightness
	opt.Channels[0].LedCount = c.LedCount
	opt.Channels[0].GpioPin = c.LedPin

	dev, err := ws2811.MakeWS2811(&opt)
	if err != nil {
		return nil, err
	}

	return &led{ws: dev}, nil
}
