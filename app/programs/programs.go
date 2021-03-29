package programs

import "github.com/brnck/go-disco/app/output"

type Programs struct {
	programs map[string]program
}

type program interface {
	run(output output.Output)
}

func New() *Programs {
	return &Programs{
		programs: make(map[string]program),
	}
}

func (p *Programs) AddProgram(name string, program program) error {
	if _, exists := p.programs[name]; exists {
		return ErrProgramExists
	}

	p.programs[name] = program

	return nil
}

func (p *Programs) RunProgram(name string, o output.Output) error {
	if _, exists := p.programs[name]; !exists {
		return ErrProgramNotExists
	}

	p.programs[name].run(o)

	return nil
}
