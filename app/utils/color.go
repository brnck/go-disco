package utils

func RgbToColor(r uint32, g uint32, b uint32) uint32 {
	return r<<16 + g<<8 + b
}
