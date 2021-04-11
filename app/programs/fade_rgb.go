package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"sync"
)

const fadeRgbName = "fade_rgb"

type fadeRgb struct {
	name string
}

func (fr fadeRgb) getName() string {
	return fr.name
}

func (fr fadeRgb) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		for j := 0; j < 256; j++ {
			if i == 0 {
				setLEDColorInRange(o, c.Start, c.End, j, 0, 0)
			}

			if i == 1 {
				setLEDColorInRange(o, c.Start, c.End, 0, j, 0)
			}

			if i == 2 {
				setLEDColorInRange(o, c.Start, c.End, 0, 0, j)
			}
			o.Render()
		}

		for j := 256; j > 0; j-- {
			if i == 0 {
				setLEDColorInRange(o, c.Start, c.End, j, 0, 0)
			}

			if i == 1 {
				setLEDColorInRange(o, c.Start, c.End, 0, j, 0)
			}

			if i == 2 {
				setLEDColorInRange(o, c.Start, c.End, 0, 0, j)
			}
			o.Render()
		}
	}
}

func NewFadeRgb() *fadeRgb {
	return &fadeRgb{
		name: fadeRgbName,
	}
}
