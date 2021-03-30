package programs

import (
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
	"time"
)

const theaterChaseName = "theater_chase"

type theaterChase struct {
	name       string
	start      int
	end        int
	iterations int
	wait       int
}

func (tc theaterChase) getName() string {
	return tc.name
}

func (tc theaterChase) Run(o output.Output, start int, end int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < tc.iterations; i++ {
		for j := 0; j < 3; j++ {
			for z := start; z < end; z += 3 {
				o.SetLed(j+z, utils.RgbToColor(128, 128, 128))
			}
			o.Render()
			time.Sleep(time.Duration(tc.wait) * time.Millisecond)
			for z := start; z < end; z += 3 {
				o.SetLed(j+z, utils.RgbToColor(0, 0, 0))
			}
		}
	}
}

func NewTheaterChase() *theaterChase {
	return &theaterChase{
		name:       theaterChaseName,
		iterations: 100,
		wait:       50,
	}
}
