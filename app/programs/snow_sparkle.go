package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"sync"
)

const snowSparkleName = "snow_sparkle"

type snowSparkle struct {
	name string
}

func (ss snowSparkle) getName() string {
	return ss.name
}

func (ss snowSparkle) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	setLEDColorInRange(o, c.Start, c.End, c.Red, c.Green, c.Blue)

	pixel := randomNumberBetween(c.Start, c.End)

	o.SetLed(pixel, utils.RgbToColor(255, 255, 255))
	o.Render()

	sleepMilliseconds(c.WaitTime)

	o.SetLed(pixel, utils.RgbToColor(c.Red, c.Green, c.Blue))
	o.Render()
}

func NewSnowSparkle() *snowSparkle {
	return &snowSparkle{
		name: snowSparkleName,
	}
}
