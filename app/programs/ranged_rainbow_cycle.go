package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
	"time"
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
	for i := 0; i < c.Iterations*128; i++ {
		for z := c.Start; z < c.End; z++ {
			o.SetLed(z, utils.ColorRoll(((z*256/c.End)+i)&255))
		}
		o.Render()
		time.Sleep(time.Duration(c.WaitTime) * time.Millisecond)
	}
}

func NewRangedRainbowCycle() *rangedRainbowCycle {
	return &rangedRainbowCycle{
		name: rangedRainbowCycleName,
	}
}
