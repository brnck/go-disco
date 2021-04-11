package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
)

const theaterChaseName = "theater_chase"

type theaterChase struct {
	name string
}

func (tc theaterChase) getName() string {
	return tc.name
}

func (tc theaterChase) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < c.Iterations; i++ {
		for j := 0; j < 3; j++ {
			for z := c.Start; z < c.End; z += 3 {
				o.SetLed(j+z, utils.RgbToColor(128, 128, 128))
			}
			o.Render()
			sleepMilliseconds(c.WaitTime)
			for z := c.Start; z < c.End; z += 3 {
				o.SetLed(j+z, utils.RgbToColor(0, 0, 0))
			}
		}
	}
}

func NewTheaterChase() *theaterChase {
	return &theaterChase{
		name: theaterChaseName,
	}
}
