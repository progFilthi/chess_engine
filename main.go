package main

import (
	"os"

	"github.com/progFilthi/chess_engine/board"
)

func main() {

	b := board.NewBoard()

	b.Print(os.Stdout)
}
