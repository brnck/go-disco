package programs

import (
	"github.com/brnck/go-disco/app/output"
	"sync"
)

type Programs struct {
	programs map[string]Program
}

type Program interface {
	getName() string
	Run(output output.Output, start int, end int, wg *sync.WaitGroup)
}

func New() *Programs {
	return &Programs{
		programs: make(map[string]Program),
	}
}

func (p *Programs) AddProgram(program Program) error {
	if _, exists := p.programs[program.getName()]; exists {
		return ErrProgramExists
	}

	p.programs[program.getName()] = program

	return nil
}

func (p *Programs) GetProgram(name string) (Program, error) {
	if _, exists := p.programs[name]; !exists {
		return nil, ErrProgramNotExists
	}

	return p.programs[name], nil
}
