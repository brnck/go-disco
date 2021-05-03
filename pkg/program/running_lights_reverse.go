package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"math"
	"sync"
)

const runningLightsReverseName = "running_lights_reverse"

type runningLightsReverse struct {
	name string
}

func NewRunningLightsReverse() *runningLightsReverse {
	return &runningLightsReverse{
		name: runningLightsReverseName,
	}
}

func (rlr runningLightsReverse) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	position := 0
	for i := c.End; i > c.Start; i-- {
		position -= 1

		for j := c.End; j > c.Start; j-- {
			ratio := (math.Sin(float64(j+position))*127 + 128) / 255

			r := int(math.Floor(ratio * float64(c.Red)))
			g := int(math.Floor(ratio * float64(c.Green)))
			b := int(math.Floor(ratio * float64(c.Blue)))

			o.SetLed(j-1, rgbToColor(r, g, b))
		}

		o.Render()

		sleepMilliseconds(c.WaitTime)
	}
}

func (rlr runningLightsReverse) getName() string {
	return rlr.name
}
