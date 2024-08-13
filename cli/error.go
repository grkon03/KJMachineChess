package cli

import "fmt"

var (
	ErrTheNameIsUsed   = fmt.Errorf("the process name is used")
	ErrProcessNotFound = fmt.Errorf("Process of the name is not found")
)
