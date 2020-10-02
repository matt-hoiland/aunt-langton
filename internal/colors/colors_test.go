package colors

import (
	"image/color"
	"testing"
)

// TestPaletteIndex tests colors.PaletteIndex
func TestPaletteIndex(t *testing.T) {
	tests := []struct {
		name     string
		data     color.Color
		expected int
	}{
		{"Unvisited", Unvisited, 0},
		{"Black", Black, 1},
		{"White", White, 2},
		{"Red", Red, 3},
	}

	for _, test := range tests {
		got := PaletteIndex(test.data)
		if got != test.expected {
			t.Errorf("PaletteIndex(%s) = %d, want %d", test.name, got, test.expected)
		}
	}

	if untested := len(ColorPalette) - len(tests); untested > 0 {
		t.Errorf("There are %d untested colors added to ColorPalette. Write tests for them!\n", untested)
	}
}
