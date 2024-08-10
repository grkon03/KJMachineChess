package cli

import (
	"bufio"
	"fmt"
	"os"
)

type CLI struct {
	Proecesses      []Process
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

func (cli *CLI) ReceiveCommand() {
	fmt.Println(splitLine)
	fmt.Println("(=°ω°=) < Meneow")
	fmt.Println(splitLine)

	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cli.ProcessingCommand = scanner.Text()
	cli.CommandsHistory = append([]string{cli.ProcessingCommand}, cli.CommandsHistory...)
}
