package utils

func RgbToColor(r int, g int, b int) uint32 {
	return uint32(r<<16 + g<<8 + b)
}

func ColorRoll(position int) uint32 {
	if position < 85 {
		return RgbToColor(position*3, 255-position*3, 0)
	}

	if position < 170 {
		position -= 85

		return RgbToColor(255-position*3, 0, position*3)
	}

	position -= 170

	return RgbToColor(0, position*3, 255-position*3)
}
