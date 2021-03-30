package programs

import (
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
	"time"
)

const rangedRainbowCycleName = "ranged_rainbow_cycle"

type rangedRainbowCycle struct {
	name       string
	start      int
	end        int
	iterations int
	wait       int
}

func (rrc rangedRainbowCycle) getName() string {
	return rrc.name
}

func (rrc rangedRainbowCycle) Run(o output.Output, start int, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < rrc.iterations*128; i++ {
		for z := start; z < end; z++ {
			o.SetLed(z, utils.ColorRoll(((z*256/end)+i)&255))
		}
		o.Render()
		time.Sleep(time.Duration(rrc.wait) * time.Millisecond)
	}
}

func NewRangedRainbowCycle() *rangedRainbowCycle {
	return &rangedRainbowCycle{
		name:       rangedRainbowCycleName,
		iterations: 50,
		wait:       10,
	}
}
