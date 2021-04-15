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

		o.SetLed(i, utils.RgbToColor(r, g, b))

		if i > 0 {
			o.SetLed(i-1, utils.RgbToColor(r, g, b))
		}

		o.Render()
		sleepMilliseconds(c.Speed)
	}

	sleepMilliseconds(c.WaitTime)

	for i := c.Start; i < c.End; i++ {
		for j := c.End - 1; j > c.Start; j-- {
			oldColor := o.GetLedColor(j - 1)
			o.SetLed(j, oldColor)
		}

		o.SetLed(i, utils.RgbToColor(0, 0, 0))
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
