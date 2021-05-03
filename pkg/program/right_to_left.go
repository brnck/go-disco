package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const rightToLeftName = "right_to_left"

type rightToLeft struct {
	name string
}

func NewRightToLeft() *rightToLeft {
	return &rightToLeft{
		name: rightToLeftName,
	}
}

func (rtl rightToLeft) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := c.End - c.Size - 2; i > c.Start; i-- {
		setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)

		for j := 1; j < c.Size+1; j++ {
			o.SetLed(i+j, rgbToColor(c.Red, c.Green, c.Blue))
		}

		o.Render()
		sleepMilliseconds(c.Speed)
	}
	sleepMilliseconds(c.WaitTime)
}

func (rtl rightToLeft) getName() string {
	return rtl.name
}
