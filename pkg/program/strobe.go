package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const strobeName = "strobe"

type strobe struct {
	name string
}

func NewStrobe() *strobe {
	return &strobe{
		name: strobeName,
	}
}

func (s strobe) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < c.Iterations; i++ {
		setLEDColorInRange(o, c.Start, c.End, c.Red, c.Green, c.Blue)
		sleepMilliseconds(c.WaitTime)
		setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)
		sleepMilliseconds(c.WaitTime)
	}

	setLEDColorInRange(o, c.Start, c.End, c.Red, c.Green, c.Blue)
}

func (s strobe) getName() string {
	return s.name
}
