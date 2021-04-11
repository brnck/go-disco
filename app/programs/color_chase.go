package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
)

const colorChaseName = "color_chase"

type colorChase struct {
	name string
}

func (cc colorChase) getName() string {
	return cc.name
}

func (cc colorChase) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := c.Start; i < c.End; i++ {
		o.SetLed(i, utils.RgbToColor(c.Red, c.Green, c.Blue))
		o.Render()

		sleepMilliseconds(c.WaitTime)

		o.SetLed(i, utils.RgbToColor(0, 0, 0))
		o.Render()
	}
}

func NewColorChase() *colorChase {
	return &colorChase{
		name: colorChaseName,
	}
}
