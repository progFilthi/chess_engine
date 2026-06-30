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

func TestStartingPosition(t *testing.T) {
	b := NewBoard()

	SetStartingPosition(&b)

	if b.Pieces[6].PieceType != Knight {
		t.Errorf("Piece on index 6 should be a Knight. (got %v)", b.Pieces[6].PieceType)
	}

	if b.Pieces[6].Color != Black {
		t.Errorf("Piece on index 6 should be a Black Knight. (got %v)", b.Pieces[6].Color)
	}

}
