package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"sync"
)

const strobeName = "strobe"

type strobe struct {
	name string
}

func (s strobe) getName() string {
	return s.name
}

func (s strobe) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < c.Iterations; i++ {
		setLEDColorInRange(o, c.Start, c.End, c.Red, c.Green, c.Blue)
		sleepMilliseconds(c.WaitTime)
		setLEDColorInRange(o, c.Start, c.End, 0, 0, 0)
		sleepMilliseconds(c.WaitTime)
	}
}

func NewStrobe() *strobe {
	return &strobe{
		name: strobeName,
	}
}
