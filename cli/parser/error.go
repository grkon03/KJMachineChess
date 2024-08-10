package pgnparser

import "fmt"

var (
	ErrResultUncaught     = fmt.Errorf("there is no result in this pgn")
	ErrUnexpectedPGNParse = fmt.Errorf("unexpected error happens at parsing this pgn")
)
