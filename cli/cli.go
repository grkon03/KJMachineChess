package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

type CLI struct {
	Processes       []Process
	CommandsHistory []string

	ProcessingCommand string
}

const launchLogo string = `
---------------------------------
KJMachine for chess
	from grkon
---------------------------------
`

const splitLine string = "---------------------"

func NewCLI() CLI {
	fmt.Print(launchLogo)
	return CLI{}
}

func (c *CLI) ReceiveCommand() {
	fmt.Println(splitLine)
	fmt.Println("(=°ω°=) < Meneow")
	fmt.Println(splitLine)

	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	c.ProcessingCommand = scanner.Text()
	c.CommandsHistory = append([]string{c.ProcessingCommand}, c.CommandsHistory...)
}

func (c *CLI) AddNewProcess(pgn string, processName string) {
	prc, err := NewProcess(pgn, processName)
	if err != nil {
		color.Red(err.Error())
	} else {
		color.Green("Successfully load PGN")
	}
	c.Processes = append(c.Processes, prc)
}
