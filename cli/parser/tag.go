package pgnparser

type TagParser struct {
	pgn   []rune
	index int
}

func (tp TagParser) Parse() (t *Tag, index int, err error) {
	return
}

func ParseTag(p *Parser) (t *Tag, index int, err error) {
	tp := TagParser{p.PGN, p.Index}
	return tp.Parse()
}
