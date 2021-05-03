package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const theaterRainbowChaseName = "theater_rainbow_chase"

type theaterRainbowChase struct {
	name string
}

func NewTheaterRainbowChase() *theaterRainbowChase {
	return &theaterRainbowChase{
		name: theaterRainbowChaseName,
	}
}

func (trc theaterRainbowChase) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < c.Iterations; i++ {
		for j := 0; j < 3; j++ {
			for k := c.Start; k < c.End; k += 3 {
				o.SetLed(j+k, colorRoll((k+i)%255))
			}
			o.Render()
			sleepMilliseconds(c.WaitTime)
			for z := c.Start; z < c.End; z += 3 {
				o.SetLed(j+z, rgbToColor(0, 0, 0))
			}
		}
	}
}

func (trc theaterRainbowChase) getName() string {
	return trc.name
}
