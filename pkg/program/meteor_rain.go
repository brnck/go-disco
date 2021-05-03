package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const meteorRainName = "meteor_rain"

type meteorRain struct {
	name string
}

func NewMeteorRain() *meteorRain {
	return &meteorRain{
		name: meteorRainName,
	}
}

func (mr meteorRain) Run(o output.Output, c Config, wg *sync.WaitGroup) {
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
				o.SetLed(i-j+c.Start, rgbToColor(c.Red, c.Green, c.Blue))
			}
		}

		o.Render()
		sleepMilliseconds(c.WaitTime)
	}
}

func (mr meteorRain) getName() string {
	return mr.name
}
