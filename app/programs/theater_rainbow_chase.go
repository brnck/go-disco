package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
)

const theaterRainbowChaseName = "theater_rainbow_chase"

type theaterRainbowChase struct {
	name string
}

func (trc theaterRainbowChase) getName() string {
	return trc.name
}

func (trc theaterRainbowChase) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < c.Iterations; i++ {
		for j := 0; j < 3; j++ {
			for k := c.Start; k < c.End; k += 3 {
				o.SetLed(j+k, utils.ColorRoll((k+i)%255))
			}
			o.Render()
			sleepMilliseconds(c.WaitTime)
			for z := c.Start; z < c.End; z += 3 {
				o.SetLed(j+z, utils.RgbToColor(0, 0, 0))
			}
		}
	}
}

func NewTheaterRainbowChase() *theaterRainbowChase {
	return &theaterRainbowChase{
		name: theaterRainbowChaseName,
	}
}
