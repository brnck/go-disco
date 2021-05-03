package program

import "errors"

var (
	ErrProgramExists    = errors.New("program already exists in the list")
	ErrProgramNotExists = errors.New("program does not exist in the programs list")
)
