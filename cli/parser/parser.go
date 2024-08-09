package parser

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
	PGN   string
	Index int
}

func (p *Parser) ExcludeComment() error {
	return nil
}

// implemented honestly, TODO: refactor to smart one
func (p *Parser) Parse() (t *Tag, m *Move, err error) {
	var first rune

	for i, r := range p.PGN[p.Index:] {
		if r != ' ' {
			p.Index += i
			first = r
		}
	}

	if first == '[' {
		t, p.Index, err = ParseTag(p)
		return
	} else {
		m, p.Index, err = ParseNotation(p)
		return
	}
}
