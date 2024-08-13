package cli

import (
	clie "github.com/grkon03/KJMachineChess/cli/engine"
	clied "github.com/grkon03/KJMachineChess/cli/engine/dataman"
	pgnparser "github.com/grkon03/KJMachineChess/cli/pgnparser"
)

type Process struct {
	PPGN        pgnparser.ParsedPGN
	Record      clied.GameRecord
	ProcessName string
	Engn        clie.Engine
}

func NewProcess(pgn string, processName string) (prc Process, err error) {
	prc.PPGN, err = pgnparser.NewParsedPGN(pgn)
	prc.ProcessName = processName

	return
}
