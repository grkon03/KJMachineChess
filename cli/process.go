package cli

import (
	cli "github.com/grkon03/KJMachineChess/cli/engine/dataman"
	pgnparser "github.com/grkon03/KJMachineChess/cli/pgnparser"
)

type Process struct {
	PPGN   pgnparser.ParsedPGN
	Record cli.GameRecord
}

func MakeProcessFromPGN(pgn string) (prc Process, err error) {
	prc.PPGN, err = pgnparser.StringToParsedPGN(pgn)

	return
}
