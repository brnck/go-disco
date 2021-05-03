package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const outsideToCenterName = "outside_to_center"

type outsideToCenter struct {
	name string
}

func NewOutsideToCenter() *outsideToCenter {
	return &outsideToCenter{
		name: outsideToCenterName,
	}
}

func (otc outsideToCenter) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	ledMax := (c.End - c.Start - c.Size/2) + c.Start
	for i := c.Start; i < ledMax; i++ {
		setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)

		for j := 1; j < c.Size+1; j++ {
			pixel := i + j
			if pixel < 0 {
				pixel = 0
			} else if pixel >= c.End {
				pixel = c.End - 1
			}
			o.SetLed(pixel, rgbToColor(c.Red, c.Green, c.Blue))

			pixel = c.Start + c.End - i - j
			if pixel < 0 {
				pixel = 0
			} else if pixel >= c.End {
				pixel = c.End - 1
			}
			o.SetLed(pixel, rgbToColor(c.Red, c.Green, c.Blue))
		}

		o.Render()
		sleepMilliseconds(c.Speed)
	}
	sleepMilliseconds(c.WaitTime)
}

func (otc outsideToCenter) getName() string {
	return otc.name
}
