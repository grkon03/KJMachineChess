package cli

import (
	clie "github.com/grkon03/KJMachineChess/cli/engine"
	clied "github.com/grkon03/KJMachineChess/cli/engine/dataman"
	pgnparser "github.com/grkon03/KJMachineChess/cli/pgnparser"
)

type EngineInterface interface {
	NewEngine() clie.Engine
	PPGNToNotativeRecord(pgnparser.ParsedPGN) clied.NotativeRecord
}
