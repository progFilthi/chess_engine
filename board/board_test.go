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

	if b.Pieces[60].PieceType != King {
		t.Errorf("Piece on index 60 should be a King. (got %v)", b.Pieces[60].PieceType)
	}

	if b.Pieces[60].Color != White {
		t.Errorf("Piece on index 60 should be a White King. (got %v)", b.Pieces[60].Color)
	}

	if b.Pieces[48].PieceType != Pawn {
		t.Errorf("Piece on index 48 should be a Pawn. (got %v)", b.Pieces[48].PieceType)
	}

	if b.Pieces[48].Color != White {
		t.Errorf("Piece on index 48 should be a White Pawn. (got %v)", b.Pieces[48].Color)
	}

	if b.Pieces[3].PieceType != Queen {
		t.Errorf("Piece on index 3 should be a Queen. (got %v)", b.Pieces[3].PieceType)
	}

	if b.Pieces[3].Color != Black {
		t.Errorf("Piece on index 3 should be a Black Queen. (got %v)", b.Pieces[3].Color)
	}

}
