package parsedcmd

import "fmt"

type ParsedCMD interface {
	Run() error
}

var (
	ErrInvalidCommand = fmt.Errorf("this command is invalid")
)

type ErrorCMD struct{}

func NewErrorCMD() ErrorCMD {
	return ErrorCMD{}
}

func (ErrorCMD) Run() error {
	return ErrInvalidCommand
}
