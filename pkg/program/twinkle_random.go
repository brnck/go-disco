package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const twinkleRandomName = "twinkle_random"

type twinkleRandom struct {
	name string
}

func NewTwinkleRandom() *twinkleRandom {
	return &twinkleRandom{
		name: twinkleRandomName,
	}
}

func (tr twinkleRandom) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)

	for i := 0; i < c.Iterations; i++ {
		pixel := randomNumberBetween(c.Start, c.End)
		r := randomNumberBetween(0, 256)
		g := randomNumberBetween(0, 256)
		b := randomNumberBetween(0, 256)

		o.SetLed(pixel, rgbToColor(r, g, b))
		o.Render()
		sleepMilliseconds(c.WaitTime)
		if c.OneLedPerScene {
			setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)
		}
	}

	sleepMilliseconds(c.WaitTime)
}

func (tr twinkleRandom) getName() string {
	return tr.name
}
