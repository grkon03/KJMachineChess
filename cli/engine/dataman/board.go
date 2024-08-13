package clied

type Color interface {
	IsWhite() bool
	IsBlack() bool
}
type PieceType interface{}
type Piece interface {
	GetColor() Color
	GetPieceType() PieceType
}
type Square interface{}
type Move interface {
	GetPiece() Piece
	GetFromSquare() Square
	GetToSquare() Square
	GetPromotionPiece() Piece
	ToNotation() string
}
type Board interface {
	DoMove(Move) (Board, error)
	GetPiece(Square) Piece
	InitializePosition()
}
