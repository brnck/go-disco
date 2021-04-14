package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
)

const staticRainbowName = "static_rainbow"

type staticRainbow struct {
	name string
}

func (sr staticRainbow) getName() string {
	return sr.name
}

func (sr staticRainbow) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 256; i++ {
		for j := c.Start; j < c.End; j++ {
			roll := (j*256/c.End + i) & 255
			o.SetLed(j, utils.ColorRoll(roll))
		}
	}
	o.Render()
	sleepMilliseconds(c.WaitTime)
}

func NewStaticRainbow() *staticRainbow {
	return &staticRainbow{
		name: staticRainbowName,
	}
}
