package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/matt-hoiland/aunt-langton/internal/ant"
	"github.com/matt-hoiland/aunt-langton/internal/board"
	"github.com/matt-hoiland/aunt-langton/internal/colors"
)

var (
	maxRows = flag.Int("r", 1028, "rows in board")
	maxCols = flag.Int("c", 1028, "columns in board")
	outFile = flag.String("o", "default.png", "output png file")
)

func main() {
	board := board.NewBoard(*maxRows, *maxCols)
	crawler := ant.NewAnt(*maxRows/2, *maxCols/2)

	for (image.Point{crawler.J(), crawler.I()}).In(board.Bounds()) {
		switch board[crawler.I()][crawler.J()] {
		case colors.PaletteIndex(colors.Unvisited):
			fallthrough
		case colors.PaletteIndex(colors.Black):
			board[crawler.I()][crawler.J()] = colors.PaletteIndex(colors.White)
			crawler.TurnLeft()
		case colors.PaletteIndex(colors.White):
			board[crawler.I()][crawler.J()] = colors.PaletteIndex(colors.Black)
			crawler.TurnRight()
		default:
			log.Fatal("Unknown color value in board, should be impossible")
		}
		crawler.Advance()
	}

	f, err := os.Create(*outFile)
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(f, board)

	fmt.Printf("STATS: steps = %10d\n", crawler.Steps())
}
