package programs

import (
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"time"
)

const theaterChaseName = "theater_chase"

type theaterChase struct {
	Name       string
	start      int
	end        int
	iterations int
	wait       int
}

func NewTheaterChase() *theaterChase {
	return &theaterChase{
		Name:       theaterChaseName,
		start:      0,
		end:        10,
		iterations: 10,
		wait:       10,
	}
}

func (tc theaterChase) run(o output.Output) {
	for i := 0; i < tc.iterations; i++ {
		for j := 0; j < 3; j++ {
			for z := tc.start; z < tc.end; z += 3 {
				o.SetLed(j+z, utils.RgbToColor(128, 128, 128))
			}
			o.Render()
			time.Sleep(time.Duration(tc.wait))
			for z := tc.start; z < tc.end; z += 3 {
				o.SetLed(j+z, utils.RgbToColor(0, 0, 0))
			}
		}
	}
}
