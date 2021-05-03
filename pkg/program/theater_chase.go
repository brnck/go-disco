package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const theaterChaseName = "theater_chase"

type theaterChase struct {
	name string
}

func NewTheaterChase() *theaterChase {
	return &theaterChase{
		name: theaterChaseName,
	}
}

func (tc theaterChase) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < c.Iterations; i++ {
		for j := 0; j < 3; j++ {
			for z := c.Start; z < c.End; z += 3 {
				o.SetLed(j+z, rgbToColor(128, 128, 128))
			}
			o.Render()
			sleepMilliseconds(c.WaitTime)
			for z := c.Start; z < c.End; z += 3 {
				o.SetLed(j+z, rgbToColor(0, 0, 0))
			}
		}
	}
}

func (tc theaterChase) getName() string {
	return tc.name
}
