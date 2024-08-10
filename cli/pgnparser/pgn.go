package pgnparser

type Result int

const (
	WhiteWon Result = iota
	BlackWon
	Drew
	NotResult
)

type ParsedPGN struct {
	Tags   []Tag
	Moves  []Move
	Result Result
}

func NewParsedPGN(pgn string) (ppgn ParsedPGN, err error) {
	psr := NewParser(pgn)
	var tag *Tag
	var move *Move
	var res Result

	err = psr.ExcludeComment()
	if err != nil {
		return
	}

	for {
		tag, move, res, err = psr.Parse()
		if tag != nil {
			ppgn.Tags = append(ppgn.Tags, *tag)
		} else if move != nil {
			ppgn.Moves = append(ppgn.Moves, *move)
		} else if res != NotResult {
			ppgn.Result = res
			break
		} else {
			if err == nil {
				err = ErrUnexpectedPGNParse
			}
			break
		}
	}

	return
}
