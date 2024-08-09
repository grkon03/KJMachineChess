package cli

type Color interface{}
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
}
type Board interface {
	DoMove(Move) (Board, error)
	GetPiece(Square) Piece
	InitializePosition()
}
