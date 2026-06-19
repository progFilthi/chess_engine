package board

import (
	"fmt"
	"io"
)

func (board *Board) Print(w io.Writer) {

	files := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	pieces := []string{"*", "P", "N", "B", "R", "Q", "K"}

	for i, piece := range board.Pieces {

		rank := i / 8

		file := i % 8

		if file == 0 {
			write(w, "%d ", 8-rank)
		}

		write(w, "%s ", pieces[piece.PieceType])

		if file == 7 {
			write(w, "\n")

		}

	}

	write(w, " ")

	for _, file := range files {
		write(w, "%s ", file)
	}
}

func write(w io.Writer, format string, args ...any) {
	_, err := fmt.Fprintf(w, format, args...)
	if err != nil {
		return
	}
}
