package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"sync"
	"time"
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
		setAllLEDColor(o, c.End, c.Red, c.Green, c.Blue)
		time.Sleep(time.Duration(c.WaitTime))
		setAllLEDColor(o, c.End, 0, 0, 0)
	}
}

func NewStrobe() *strobe {
	return &strobe{
		name: strobeName,
	}
}
