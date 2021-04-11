package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
)

const colorChaseReverseName = "color_chase_reverse"

type colorChaseReverse struct {
	name string
}

func (ccr colorChaseReverse) getName() string {
	return ccr.name
}

func (ccr colorChaseReverse) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := c.End; i > c.Start; i-- {
		o.SetLed(i, utils.RgbToColor(c.Red, c.Green, c.Blue))
		o.Render()

		sleepMilliseconds(c.WaitTime)

		o.SetLed(i, utils.RgbToColor(0, 0, 0))
		o.Render()
	}
}

func NewColorChaseReverse() *colorChaseReverse {
	return &colorChaseReverse{
		name: colorChaseReverseName,
	}
}
