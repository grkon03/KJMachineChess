package pgnparser

func ParseResult(p *Parser) (Result, error) {
	var SpecificChars []rune

	for _, r := range p.PGN[p.Index:] {
		if r != ' ' {
			SpecificChars = append(SpecificChars, r)
		}
	}

	switch string(SpecificChars) {
	case "1-0":
		return WhiteWon, nil
	case "0-1":
		return BlackWon, nil
	case "1/2-1/2":
		return Drew, nil
	default:
		return NotResult, ErrResultUncaught
	}
}
