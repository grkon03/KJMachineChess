package cli

import cli "github.com/grkon03/KJMachineChess/cli/engine/dataman"

type RecErrMessage string

type RecErr interface {
	GetTurn() (int, cli.Color)
	What() RecErrMessage
}
