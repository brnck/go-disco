package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"math"
	"sync"
)

const runningLightsName = "running_lights"

type runningLights struct {
	name string
}

func (rl runningLights) getName() string {
	return rl.name
}

func (rl runningLights) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	position := 0
	for i := c.Start; i < c.End*2; i++ {
		position += 1

		for j := c.Start; j < c.End; j++ {
			ratio := (math.Sin(float64(j+position))*127 + 128) / 255

			r := int(math.Floor(ratio * float64(c.Red)))
			g := int(math.Floor(ratio * float64(c.Green)))
			b := int(math.Floor(ratio * float64(c.Blue)))

			o.SetLed(j, utils.RgbToColor(r, g, b))
		}

		o.Render()

		sleepMilliseconds(c.WaitTime)
	}
}

func NewRunningLights() *runningLights {
	return &runningLights{
		name: runningLightsName,
	}
}
