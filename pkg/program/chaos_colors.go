package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const chaosColorsName = "chaos_colors"

type chaosColors struct {
	name string
}

func NewChaosColors() *chaosColors {
	return &chaosColors{
		name: chaosColorsName,
	}
}

func (cc chaosColors) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := c.Start; i < c.End; i++ {
		r := randomNumberBetween(0, 256)
		g := randomNumberBetween(0, 256)
		b := randomNumberBetween(0, 256)

		o.SetLed(i, rgbToColor(r, g, b))
	}
	o.Render()
	sleepMilliseconds(c.WaitTime)
}

func (cc chaosColors) getName() string {
	return cc.name
}
