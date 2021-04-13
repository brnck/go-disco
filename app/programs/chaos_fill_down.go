package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
)

const chaosFillDownName = "chaos_fill_down"

type chaosFillDown struct {
	name string
}

func (cfd chaosFillDown) getName() string {
	return cfd.name
}

func (cfd chaosFillDown) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)

	for i := c.Start; i < c.End; i++ {
		r := randomNumberBetween(0, 256)
		g := randomNumberBetween(0, 256)
		b := randomNumberBetween(0, 256)

		for j := c.Start; j < c.End-i; j++ {
			o.SetLed(j, utils.RgbToColor(r, g, b))

			if j > 0 {
				o.SetLed(j-1, utils.RgbToColor(r, g, b))
			}

			sleepMilliseconds(c.Speed)
		}
		o.Render()
	}

	sleepMilliseconds(c.WaitTime)

	for i := c.End - 1; i > c.Start; i-- {
		for j := c.End; i > c.Start; j-- {
			oldColor := o.GetLedColor(j - 1)
			o.SetLed(j, oldColor)
		}

		o.SetLed(i-c.End+1, utils.RgbToColor(0, 0, 0))
		o.Render()
		sleepMilliseconds(c.Speed)
	}

	sleepMilliseconds(c.WaitTime)
}

func NewChaosFillDown() *chaosFillDown {
	return &chaosFillDown{
		name: chaosFillDownName,
	}
}
