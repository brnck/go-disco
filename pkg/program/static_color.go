package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const staticColorName = "static_color"

type staticColor struct {
	name string
}

func NewStaticColor() *staticColor {
	return &staticColor{
		name: staticColorName,
	}
}

func (sc staticColor) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	setLEDColorInRange(o, c.Start, c.End, c.Red, c.Green, c.Blue)
}

func (sc staticColor) getName() string {
	return sc.name
}
