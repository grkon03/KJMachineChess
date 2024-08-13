package parsedcmd

import "github.com/grkon03/KJMachineChess/cli"

/*
Regist PGN to cli

CMD:
regist [process name] [pgn]
regist [process name] -f [file name]
*/
type PsdRegistPGN struct {
	c           *cli.CLI
	pgn         string
	processName string
}

func (r PsdRegistPGN) Run() error {
	for _, p := range r.c.Processes {
		if p.ProcessName == r.processName {
			return cli.ErrTheNameIsUsed
		}
	}
	r.c.AddNewProcess(r.pgn, r.processName)
	return nil
}
