package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"math"
	"sync"
)

const runningLightsReverseName = "running_lights_reverse"

type runningLightsReverse struct {
	name string
}

func (rlr runningLightsReverse) getName() string {
	return rlr.name
}

func (rlr runningLightsReverse) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	position := 0
	for i := c.End * 2; i > c.Start; i-- {
		position -= 1

		for j := c.End; j > c.Start; j-- {
			ratio := (math.Sin(float64(j+position))*127 + 128) / 255

			r := int(math.Floor(ratio * float64(c.Red)))
			g := int(math.Floor(ratio * float64(c.Green)))
			b := int(math.Floor(ratio * float64(c.Blue)))

			o.SetLed(j-1, utils.RgbToColor(r, g, b))
		}

		o.Render()

		sleepMilliseconds(c.WaitTime)
	}
}

func NewRunningLightsReverse() *runningLightsReverse {
	return &runningLightsReverse{
		name: runningLightsReverseName,
	}
}
