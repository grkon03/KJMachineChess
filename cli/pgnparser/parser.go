package pgnparser

import cli "github.com/grkon03/KJMachineChess/cli/engine/dataman"

type Tag struct {
	Key   string
	Value string
}

type Move struct {
	Turn     int
	Color    cli.Color
	TheMove  cli.Move
	Appendix []rune
}

type Parser struct {
	PGN   []rune
	Index int
}

func NewParser(pgn string) Parser {
	return Parser{[]rune(pgn), 0}
}

func (p *Parser) ExcludeComment() error {
	return nil
}

// implemented honestly, TODO: refactor to smart one
func (p *Parser) Parse() (tag *Tag, move *Move, res Result, err error) {
	var first rune
	tag = nil
	move = nil
	res = NotResult

	for i, r := range p.PGN[p.Index:] {
		if r != ' ' {
			p.Index += i
			first = r
		}
	}

	if first == '[' {
		tag, p.Index, err = ParseTag(p)
		return
	} else {
		move, p.Index, err = ParseNotation(p)

		if move == nil {
			res, err = ParseResult(p)
		}
		return
	}
}
