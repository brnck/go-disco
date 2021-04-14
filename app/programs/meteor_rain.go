package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
)

const meteorRainName = "meteor_rain"

type meteorRain struct {
	name string
}

func (mr meteorRain) getName() string {
	return mr.name
}

func (mr meteorRain) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)
	for i := c.Start; i < c.End; i++ {
		for j := c.Start; j < c.End; j++ {
			if !c.RandomDecay || randomNumberBetween(0, 10) > 5 {
				fadeToBlack(o, j, uint32(c.TrailDecay))
			}
		}

		for j := c.Start; j < c.Start+c.Size; j++ {
			if (i-j+c.Start) < c.End && (i-j) >= 0 {
				o.SetLed(i-j+c.Start, utils.RgbToColor(c.Red, c.Green, c.Blue))
			}
		}

		o.Render()
		sleepMilliseconds(c.WaitTime)
	}
}

func NewMeteorRain() *meteorRain {
	return &meteorRain{
		name: meteorRainName,
	}
}
