package dto

type Color string
type PieceType string

const (
	White Color = "white"
	Black Color = "black"
)

const (
	Pawn   PieceType = "P"
	Rook   PieceType = "R"
	Knight PieceType = "N"
	Bishop PieceType = "B"
	Queen  PieceType = "Q"
	King   PieceType = "K"
)

type Piece struct {
	Type  PieceType
	Color Color
}
