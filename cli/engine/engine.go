package clie

import clied "github.com/grkon03/KJMachineChess/cli/engine/dataman"

type Engine interface {
	RegistRecord(clied.NotativeRecord)

	Run() RecordIncorrect
}
