package output

import (
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

type led struct {
	ws *ws2811.WS2811
}

func (l *led) SetLed(i int, color uint32) {
	l.ws.Leds(0)[i] = color
}

func (l *led) GetLedColor(i int) uint32 {
	return l.ws.Leds(0)[i]
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

func newLed(c *ws2811.Option) (*led, error) {
	dev, err := ws2811.MakeWS2811(c)
	if err != nil {
		return nil, err
	}

	return &led{ws: dev}, nil
}
