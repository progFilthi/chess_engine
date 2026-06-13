package board

type Color int
type PieceType int

const (
	White Color = iota
	Black
)

const (
	None PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	Color     Color
	PieceType PieceType
}

type Board struct {
	Pieces [64]Piece
}
