package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const snowSparkleName = "snow_sparkle"

type snowSparkle struct {
	name string
}

func NewSnowSparkle() *snowSparkle {
	return &snowSparkle{
		name: snowSparkleName,
	}
}

func (ss snowSparkle) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	setLEDColorInRange(o, c.Start, c.End, c.Red, c.Green, c.Blue)

	pixel := randomNumberBetween(c.Start, c.End)

	o.SetLed(pixel, rgbToColor(255, 255, 255))
	o.Render()

	sleepMilliseconds(c.WaitTime)

	o.SetLed(pixel, rgbToColor(c.Red, c.Green, c.Blue))
	o.Render()
}

func (ss snowSparkle) getName() string {
	return ss.name
}
