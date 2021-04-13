package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
)

const twinkleName = "twinkle"

type twinkle struct {
	name string
}

func (t twinkle) getName() string {
	return t.name
}

func (t twinkle) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)

	for i := 0; i < c.Iterations; i++ {
		pixel := randomNumberBetween(c.Start, c.End)

		o.SetLed(pixel, utils.RgbToColor(c.Red, c.Green, c.Blue))
		o.Render()
		sleepMilliseconds(c.WaitTime)
		if c.OneLedPerScene {
			setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)
		}
	}

	sleepMilliseconds(c.WaitTime)
}

func NewTwinkle() *twinkle {
	return &twinkle{
		name: twinkleName,
	}
}
