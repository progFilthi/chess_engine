package main

import (
	"os"

	"github.com/progFilthi/chess_engine/board"
)

func main() {

	b := board.NewBoard()

	board.SetStartingPosition(&b)

	b.Print(os.Stdout)
}
