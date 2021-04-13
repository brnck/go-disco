package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
)

const chaosColorsName = "chaos_colors"

type chaosColors struct {
	name string
}

func (cc chaosColors) getName() string {
	return cc.name
}

func (cc chaosColors) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := c.Start; i < c.End; i++ {
		r := randomNumberBetween(0, 256)
		g := randomNumberBetween(0, 256)
		b := randomNumberBetween(0, 256)

		o.SetLed(i, utils.RgbToColor(r, g, b))
	}
	o.Render()
	sleepMilliseconds(c.WaitTime)
}

func NewChaosColors() *chaosColors {
	return &chaosColors{
		name: chaosColorsName,
	}
}
