package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const sparkleName = "sparkle"

type sparkle struct {
	name string
}

func NewSparkle() *sparkle {
	return &sparkle{
		name: sparkleName,
	}
}

func (s sparkle) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)
	pixel := randomNumberBetween(c.Start, c.End)

	o.SetLed(pixel, rgbToColor(c.Red, c.Green, c.Blue))
	o.Render()

	sleepMilliseconds(c.WaitTime)

	o.SetLed(pixel, rgbToColor(0, 0, 0))
	o.Render()
}

func (s sparkle) getName() string {
	return s.name
}
