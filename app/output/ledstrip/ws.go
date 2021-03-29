package ledstrip

import (
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

type Led struct {
	ws *ws2811.WS2811
}

func (l *Led) SetLed(i int, color uint32) {
	l.ws.Leds(0)[i] = color
}

func (l *Led) Render() error {
	return l.ws.Render()
}

func (l *Led) Open() error {
	return l.ws.Init()
}

func (l *Led) Close() {
	l.ws.Fini()
}

func Init() (*Led, error) {
	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = 128
	opt.Channels[0].LedCount = 60

	opt.Channels[0] = ws2811.ChannelOption{
		GpioPin:    18,
		Invert:     false,
		LedCount:   70,
		StripeType: ws2811.WS2812Strip,
		Brightness: 64,
		WShift:     1,
		RShift:     1,
		GShift:     1,
		BShift:     1,
		Gamma:      nil,
	}

	dev, err := ws2811.MakeWS2811(&opt)
	if err != nil {
		return nil, err
	}

	return &Led{ws: dev}, nil
}
