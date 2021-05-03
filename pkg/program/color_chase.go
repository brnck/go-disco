package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const colorChaseName = "color_chase"

type colorChase struct {
	name string
}

func NewColorChase() *colorChase {
	return &colorChase{
		name: colorChaseName,
	}
}

func (cc colorChase) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := c.Start; i < c.End; i++ {
		o.SetLed(i, rgbToColor(c.Red, c.Green, c.Blue))
		o.Render()

		sleepMilliseconds(c.WaitTime)

		o.SetLed(i, rgbToColor(0, 0, 0))
		o.Render()
	}
}

func (cc colorChase) getName() string {
	return cc.name
}
