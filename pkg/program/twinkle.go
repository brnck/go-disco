package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const twinkleName = "twinkle"

type twinkle struct {
	name string
}

func NewTwinkle() *twinkle {
	return &twinkle{
		name: twinkleName,
	}
}

func (t twinkle) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)

	for i := 0; i < c.Iterations; i++ {
		pixel := randomNumberBetween(c.Start, c.End)

		o.SetLed(pixel, rgbToColor(c.Red, c.Green, c.Blue))
		o.Render()
		sleepMilliseconds(c.WaitTime)
		if c.OneLedPerScene {
			setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)
		}
	}

	sleepMilliseconds(c.WaitTime)
}

func (t twinkle) getName() string {
	return t.name
}
