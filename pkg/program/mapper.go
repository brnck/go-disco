package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const (
	mapperName = "mapper"
)

var colors = [...][3]int{
	{255, 255, 255},
	{255, 0, 0},
	{0, 255, 0},
	{0, 0, 255},
	{255, 255, 0},
	{0, 255, 255},
	{255, 0, 255},
	{128, 0, 0},
	{128, 128, 0},
	{0, 128, 0},
}

type mapper struct {
	name string
}

func NewMapper() *mapper {
	return &mapper{
		name: mapperName,
	}
}

func (m mapper) Run(o output.Output, c Config, wg *sync.WaitGroup) {
	defer wg.Done()

	for index, mapper := range c.Map {
		for i := mapper.Min; i < mapper.Max; i++ {
			o.SetLed(i, rgbToColor(colors[index][0], colors[index][1], colors[index][2]))
		}
	}
	o.Render()
	sleepMilliseconds(c.WaitTime)
}

func (m mapper) getName() string {
	return m.name
}
