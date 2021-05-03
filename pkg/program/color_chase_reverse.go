package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const colorChaseReverseName = "color_chase_reverse"

type colorChaseReverse struct {
	name string
}

func NewColorChaseReverse() *colorChaseReverse {
	return &colorChaseReverse{
		name: colorChaseReverseName,
	}
}

func (ccr colorChaseReverse) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := c.End; i > c.Start; i-- {
		o.SetLed(i, rgbToColor(c.Red, c.Green, c.Blue))
		o.Render()

		sleepMilliseconds(c.WaitTime)

		o.SetLed(i, rgbToColor(0, 0, 0))
		o.Render()
	}
}

func (ccr colorChaseReverse) getName() string {
	return ccr.name
}
