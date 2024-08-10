package cmdparser

import "github.com/grkon03/KJMachineChess/cli/cmdparser/parsedcmd"

type CMDParser struct {
	cmd   []rune
	index int
}

func NewCMDParser(cmd string) CMDParser {
	return CMDParser{[]rune(cmd), 0}
}

func (p *CMDParser) Parse() parsedcmd.ParsedCMD {
	return parsedcmd.NewErrorCMD()
}
