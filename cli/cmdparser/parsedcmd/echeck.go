package parsedcmd

import (
	"github.com/grkon03/KJMachineChess/cli"
	clie "github.com/grkon03/KJMachineChess/cli/engine"
)

/*
Check error of PGN as ignoring the appendix

CMD:
echeck [process name]
*/
type PsdECheckPGN struct {
	c           *cli.CLI
	processName string
	process     *cli.Process
}

func (ec *PsdECheckPGN) Run() error {
	if ec.process == nil {
		ec.process = ec.c.SearchProcess(ec.processName)
		if ec.process == nil {
			return cli.ErrProcessNotFound
		}
	}

	ec.process.Engn.RegistRecord(ec.c.EI.PPGNToNotativeRecord(ec.process.PPGN))
	ri := ec.process.Engn.Run()

	where, err := clie.WhereToString(ri.Where())

	if err != nil {
		panic(err)
	}

	if ri.IsNotDetected() {
		ec.c.Println("There is no illegal move.")
		return nil
	}

	ec.c.Println(where + ri.What())

	return nil
}
