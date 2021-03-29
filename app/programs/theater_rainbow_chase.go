package programs

import (
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"time"
)

const theaterRainbowChaseName = "theater_rainbow_chase"

type theaterRainbowChase struct {
	Name       string
	start      int
	end        int
	iterations int
	wait       int
}

func NewTheaterRainbowChase() *theaterRainbowChase {
	return &theaterRainbowChase{
		Name:       theaterRainbowChaseName,
		start:      0,
		end:        10,
		iterations: 10,
		wait:       10,
	}
}

func (trc theaterRainbowChase) run(o output.Output) {
	for i := 0; i < trc.iterations; i++ {
		for j := 0; j < 3; j++ {
			for k := trc.start; k < trc.end; k += 3 {
				o.SetLed(j+k, utils.ColorRoll((k+i)%255))
			}
			o.Render()
			time.Sleep(time.Duration(trc.wait))
			for z := trc.start; z < trc.end; z += 3 {
				o.SetLed(j+z, utils.RgbToColor(0, 0, 0))
			}
		}
	}
}
