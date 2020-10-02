package colors

import "image/color"

// The set of valid colors.
var (
	Unvisited = color.RGBA{0x80, 0x80, 0x80, 0xff}
	Black     = color.RGBA{0x00, 0x00, 0x00, 0xff}
	White     = color.RGBA{0xff, 0xff, 0xff, 0xff}
	Red       = color.RGBA{0xff, 0x00, 0x00, 0xff}
)

// ColorPalette is the pallet of all possible colors.
var ColorPalette = color.Palette{
	Unvisited,
	Black,
	White,
	Red,
}

// PaletteIndex returns the index of the given color in ColorPalette
func PaletteIndex(c color.Color) int {
	for i := 0; i < len(ColorPalette); i++ {
		if ColorPalette[i] == c {
			return i
		}
	}
	return -1
}
