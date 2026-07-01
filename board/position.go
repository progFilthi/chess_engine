package board

func SetStartingPosition(b *Board) {

	//black pieces
	b.Pieces[0] = Piece{Black, Rook}
	b.Pieces[1] = Piece{Black, Knight}
	b.Pieces[2] = Piece{Black, Bishop}
	b.Pieces[3] = Piece{Black, Queen}
	b.Pieces[4] = Piece{Black, King}
	b.Pieces[5] = Piece{Black, Bishop}
	b.Pieces[6] = Piece{Black, Knight}
	b.Pieces[7] = Piece{Black, Rook}

	//black pawns
	b.Pieces[8] = Piece{Black, Pawn}
	b.Pieces[9] = Piece{Black, Pawn}
	b.Pieces[10] = Piece{Black, Pawn}
	b.Pieces[11] = Piece{Black, Pawn}
	b.Pieces[12] = Piece{Black, Pawn}
	b.Pieces[13] = Piece{Black, Pawn}
	b.Pieces[14] = Piece{Black, Pawn}
	b.Pieces[15] = Piece{Black, Pawn}

	//white pieces
	b.Pieces[56] = Piece{White, Rook}
	b.Pieces[57] = Piece{White, Knight}
	b.Pieces[58] = Piece{White, Bishop}
	b.Pieces[59] = Piece{White, Queen}
	b.Pieces[60] = Piece{White, King}
	b.Pieces[61] = Piece{White, Bishop}
	b.Pieces[62] = Piece{White, Knight}
	b.Pieces[63] = Piece{White, Rook}

	//white pawns
	b.Pieces[48] = Piece{White, Pawn}
	b.Pieces[49] = Piece{White, Pawn}
	b.Pieces[50] = Piece{White, Pawn}
	b.Pieces[51] = Piece{White, Pawn}
	b.Pieces[52] = Piece{White, Pawn}
	b.Pieces[53] = Piece{White, Pawn}
	b.Pieces[54] = Piece{White, Pawn}
	b.Pieces[55] = Piece{White, Pawn}

}
