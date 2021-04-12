package programs

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"math/rand"
	"sync"
	"time"
)

const sparkleName = "sparkle"

type sparkle struct {
	name string
}

func (s sparkle) getName() string {
	return s.name
}

func (s sparkle) Run(o output.Output, c config.Program, wg *sync.WaitGroup) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())
	pixel := rand.Intn(c.End-c.Start) + c.Start

	o.SetLed(pixel, utils.RgbToColor(c.Red, c.Green, c.Blue))
	o.Render()

	sleepMilliseconds(c.WaitTime)
	o.SetLed(pixel, utils.RgbToColor(0, 0, 0))
	o.Render()
}

func NewSparkle() *sparkle {
	return &sparkle{
		name: sparkleName,
	}
}
