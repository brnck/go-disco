package programs

import (
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/utils"
	"time"
)

func setAllLEDColor(output output.Output, ledCount int, r int, g int, b int) {
	for i := 0; i < ledCount; i++ {
		output.SetLed(i, utils.RgbToColor(r, g, b))
	}
	output.Render()
}

func setLEDColorInRange(output output.Output, start int, end int, r int, g int, b int) {
	for i := start; i < end; i++ {
		output.SetLed(i, utils.RgbToColor(r, g, b))
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

	output.SetLed(position, utils.RgbToColor(int(r), int(g), int(b)))
}

func getCurrentMilis() int {
	return time.Now().Second()
}

func sleepMilliseconds(duration int) {
	time.Sleep(time.Duration(duration) * time.Millisecond)
}
