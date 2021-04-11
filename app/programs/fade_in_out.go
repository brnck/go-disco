package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"math"
	"sync"
)

const fadeInOutName = "fade_in_out"

type fadeInOut struct {
	name string
}

func (fio fadeInOut) getName() string {
	return fio.name
}

func (fio fadeInOut) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 256; i++ {
		ratio := int(math.Floor(float64(i) / 256.0))
		r := ratio * c.Red
		g := ratio * c.Green
		b := ratio * c.Blue

		setLEDColorInRange(o, c.Start, c.End, r, g, b)
		o.Render()
	}

	for i := 256; i > 0; i-- {
		ratio := int(math.Floor(float64(i) / 256.0))
		r := ratio * c.Red
		g := ratio * c.Green
		b := ratio * c.Blue

		setLEDColorInRange(o, c.Start, c.End, r, g, b)
		o.Render()
	}
}

func NewFadeInOut() *fadeInOut {
	return &fadeInOut{
		name: fadeInOutName,
	}
}
