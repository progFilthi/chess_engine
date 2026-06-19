package board

import "testing"

func TestBoard(t *testing.T) {

	b := NewBoard()

	for i, piece := range b.Pieces {

		if piece.PieceType != None {
			t.Errorf("Square %d should be empty, got %v", i, piece.PieceType)
		}
	}

}
