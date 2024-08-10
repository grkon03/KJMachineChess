package pgnparser

type NotationParser struct {
	pgn   []rune
	index int
}

func (np *NotationParser) Parse() (t *Move, index int, err error) {
	return
}

func ParseNotation(p *Parser) (t *Move, index int, err error) {
	np := NotationParser{p.PGN, p.Index}
	return np.Parse()
}
