package cli

import cli "github.com/grkon03/KJMachineChess/cli/engine/dataman"

type Engine interface {
	RegistRecord(cli.NotativeRecord)

	Run() RecErr
}
