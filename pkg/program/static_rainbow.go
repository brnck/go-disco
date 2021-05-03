package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const staticRainbowName = "static_rainbow"

type staticRainbow struct {
	name string
}

func NewStaticRainbow() *staticRainbow {
	return &staticRainbow{
		name: staticRainbowName,
	}
}

func (sr staticRainbow) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 256; i++ {
		for j := c.Start; j < c.End; j++ {
			roll := (j*256/c.End + i) & 255
			o.SetLed(j, colorRoll(roll))
		}
	}
	o.Render()
	sleepMilliseconds(c.WaitTime)
}

func (sr staticRainbow) getName() string {
	return sr.name
}
