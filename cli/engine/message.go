package clie

import (
	"fmt"

	clied "github.com/grkon03/KJMachineChess/cli/engine/dataman"
)

type RecordIncorrect interface {
	IsNotDetected() bool
	Where() (int, clied.Color)
	What() string
}

func WhereToString(i int, c clied.Color) (string, error) {
	if c.IsWhite() {
		return fmt.Sprintf("%d. ", i), nil
	} else if c.IsBlack() {
		return fmt.Sprintf("%d, ... ", i), nil
	} else {
		return "", fmt.Errorf("something wrong in detecting color")
	}
}
