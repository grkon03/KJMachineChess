package parsedcmd

import "github.com/grkon03/KJMachineChess/cli"

type PsdRegistPGN struct {
	c           *cli.CLI
	pgn         string
	processName string
}

func (r PsdRegistPGN) Run() (err error) {
	r.c.AddNewProcess(r.pgn, r.processName)
	return
}
