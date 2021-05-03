package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"math"
	"sync"
)

const fireName = "fire"

type fire struct {
	name string
	heat [][]int
}

func NewFire() *fire {
	return &fire{
		name: fireName,
		heat: make([][]int, 10),
	}
}

func (f fire) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()
	defer f.clearHeatPosition(c)

	f.initializeHeats(c)
	for i := 0; i < c.Iterations; i++ {
		f.setUpCoolDown(c)
		f.setUpSparking(c)
		for j := c.Start; j < c.End; j++ {
			o.SetLed(j, calculatePixelHeatColor(f.heat[c.HeatKey][j]))
		}

		o.Render()
		sleepMilliseconds(c.WaitTime)
	}
}

func (f fire) setUpCoolDown(c Config) {
	for i := c.Start; i < c.End; i++ {
		coolingTop := c.Start + int(math.Floor(float64(c.Cooling*10/c.End))+2)
		coolDown := randomNumberBetween(c.Start, coolingTop)
		if coolDown > f.heat[c.HeatKey][i] {
			f.heat[c.HeatKey][i] = 0
		} else {
			f.heat[c.HeatKey][i] = f.heat[c.HeatKey][i] + c.Start - coolDown
		}
	}

	for i := c.End - 1; i > c.Start+2; i-- {
		previousHeat := f.heat[c.HeatKey][i-1] + f.heat[c.HeatKey][i-2] + f.heat[c.HeatKey][i-2]
		f.heat[c.HeatKey][i] = int(math.Floor(float64(previousHeat) / 3))
	}
}

func (f fire) clearHeatPosition(c Config) {
	f.heat[c.HeatKey] = make([]int, 10)
}

func (f fire) setUpSparking(c Config) {
	if randomNumberBetween(0, 255) < c.Sparking {
		rand := randomNumberBetween(c.Start, c.Start+7)
		f.heat[c.HeatKey][rand] = f.heat[c.HeatKey][rand] + randomNumberBetween(160, 255)
		if f.heat[c.HeatKey][rand] > 255 {
			f.heat[c.HeatKey][rand] = 255
		}
	}
}

func (f fire) initializeHeats(c Config) {
	f.heat[c.HeatKey] = make([]int, c.End)
}

func calculatePixelHeatColor(heat int) uint32 {
	// Scale 'heat' down from 0-255 to 0-191
	descaledTemperature := int(math.Floor(float64(heat) / 255 * 191))
	heatRamp := (descaledTemperature & 0x3F) << 2

	color := rgbToColor(heatRamp, 0, 0)
	if descaledTemperature > 0x80 {
		color = rgbToColor(255, 255, heatRamp)
	} else if descaledTemperature > 0x40 {
		color = rgbToColor(255, heatRamp, 0)
	}

	return color
}

func (f fire) getName() string {
	return f.name
}
