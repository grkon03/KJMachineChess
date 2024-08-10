package cli

import (
	cli "github.com/grkon03/KJMachineChess/cli/engine/dataman"
	pgnparser "github.com/grkon03/KJMachineChess/cli/pgnparser"
)

type Process struct {
	PPGN        pgnparser.ParsedPGN
	Record      cli.GameRecord
	ProcessName string
}

func NewProcess(pgn string, processName string) (prc Process, err error) {
	prc.PPGN, err = pgnparser.NewParsedPGN(pgn)
	prc.ProcessName = processName

	return
}
