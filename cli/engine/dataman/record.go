package cli

type NotativeRecord interface {
	Rewind()
	Next() Move
	AddMoveBehind(Move)
}

type GameRecord interface {
	Rewind()
	Next() Board
	AddMoveBehind(Move) error

	ConvertFromNotation(NotativeRecord) error
}
