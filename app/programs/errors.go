package programs

import "errors"

var (
	ErrProgramExists    = errors.New("Program already exists in the list")
	ErrProgramNotExists = errors.New("Program does not exist in the programs list")
)
