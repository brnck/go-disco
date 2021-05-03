package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"math/rand"
	"time"
)

func setLEDColorInRange(output output.Output, start int, end int, r int, g int, b int) {
	for i := start; i < end; i++ {
		output.SetLed(i, rgbToColor(r, g, b))
	}
	output.Render()
}

func fadeToBlack(output output.Output, position int, fadeValue uint32) {
	oldColor := output.GetLedColor(position)
	r := (oldColor & 0x00ff0000) >> 16
	g := (oldColor & 0x0000ff00) >> 8
	b := oldColor & 0x000000ff

	if r <= 10 {
		r = 0
	} else {
		r = r - (r * fadeValue / 256)
	}

	if g <= 10 {
		g = 0
	} else {
		g = g - (g * fadeValue / 256)
	}

	if b <= 10 {
		b = 0
	} else {
		b = b - (b * fadeValue / 256)
	}

	output.SetLed(position, rgbToColor(int(r), int(g), int(b)))
}

func sleepMilliseconds(duration int) {
	time.Sleep(time.Duration(duration) * time.Millisecond)
}

func randomNumberBetween(start int, end int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(end-start) + start
}
