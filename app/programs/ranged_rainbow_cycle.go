package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
)

const rangedRainbowCycleName = "ranged_rainbow_cycle"

type rangedRainbowCycle struct {
	name string
}

func (rrc rangedRainbowCycle) getName() string {
	return rrc.name
}

func (rrc rangedRainbowCycle) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < c.Iterations; i++ {
		for z := c.Start; z < c.End; z++ {
			position := ((z * 256 / c.End) + i) & 255
			o.SetLed(z, utils.ColorRoll(position))
		}
		o.Render()
		sleepMilliseconds(c.WaitTime)
	}
}

func NewRangedRainbowCycle() *rangedRainbowCycle {
	return &rangedRainbowCycle{
		name: rangedRainbowCycleName,
	}
}
