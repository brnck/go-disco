package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"sync"
)

const staticColorName = "static_color"

type staticColor struct {
	name string
}

func (sc staticColor) getName() string {
	return sc.name
}

func (sc staticColor) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	setLEDColorInRange(o, c.Start, c.End, c.Red, c.Green, c.Blue)
}

func NewStaticColor() *staticColor {
	return &staticColor{
		name: staticColorName,
	}
}
