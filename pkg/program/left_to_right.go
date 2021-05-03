package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const leftToRightName = "left_to_right"

type leftToRight struct {
	name string
}

func NewLeftToRight() *leftToRight {
	return &leftToRight{
		name: leftToRightName,
	}
}

func (ltr leftToRight) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := c.Start; i < c.End-c.Size-2; i++ {
		setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)

		for j := 1; j < c.Size+1; j++ {
			o.SetLed(i+j, rgbToColor(c.Red, c.Green, c.Blue))
		}

		o.Render()
		sleepMilliseconds(c.Speed)
	}
	sleepMilliseconds(c.WaitTime)
}

func (ltr leftToRight) getName() string {
	return ltr.name
}
