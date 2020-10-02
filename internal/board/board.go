package board

import (
	"image"
	"image/color"

	"github.com/matt-hoiland/aunt-langton/internal/colors"
)

// Board represents the gameboard and satisfies image.Image interface
type Board [][]int

// NewBoard creates a new board of the given dimensions
func NewBoard(nRows, nCols int) Board {
	b := make([][]int, nRows)
	for i := 0; i < nRows; i++ {
		b[i] = make([]int, nCols)
	}
	return b
}

// ColorModel satisfies image.Image.ColorModel
func (b Board) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds satisfies image.Image.Bounds
func (b Board) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(b[0]), len(b))
}

// At satisfies image.Image.At
func (b Board) At(x, y int) color.Color {
	return colors.ColorPalette[b[y][x]]
}
