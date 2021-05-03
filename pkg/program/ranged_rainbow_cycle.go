package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const rangedRainbowCycleName = "ranged_rainbow_cycle"

type rangedRainbowCycle struct {
	name string
}

func NewRangedRainbowCycle() *rangedRainbowCycle {
	return &rangedRainbowCycle{
		name: rangedRainbowCycleName,
	}
}

func (rrc rangedRainbowCycle) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < c.Iterations; i++ {
		for z := c.Start; z < c.End; z++ {
			position := ((z * 256 / c.End) + i) & 255
			o.SetLed(z, colorRoll(position))
		}
		o.Render()
		sleepMilliseconds(c.WaitTime)
	}
}

func (rrc rangedRainbowCycle) getName() string {
	return rrc.name
}
