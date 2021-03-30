package programs

import (
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
	"time"
)

const theaterRainbowChaseName = "theater_rainbow_chase"

type theaterRainbowChase struct {
	name       string
	start      int
	end        int
	iterations int
	wait       int
}

func (trc theaterRainbowChase) getName() string {
	return trc.name
}

func (trc theaterRainbowChase) Run(o output.Output, start int, end int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < trc.iterations; i++ {
		for j := 0; j < 3; j++ {
			for k := start; k < end; k += 3 {
				o.SetLed(j+k, utils.ColorRoll((k+i)%255))
			}
			o.Render()
			time.Sleep(time.Duration(trc.wait) * time.Millisecond)
			for z := start; z < end; z += 3 {
				o.SetLed(j+z, utils.RgbToColor(0, 0, 0))
			}
		}
	}
}

func NewTheaterRainbowChase() *theaterRainbowChase {
	return &theaterRainbowChase{
		name:       theaterRainbowChaseName,
		iterations: 100,
		wait:       50,
	}
}
