package scene

import (
	"github.com/brnck/go-disco/pkg/program"
)

type SceneProgram struct {
	programRunner program.Program
	programConfig program.Config
}

func NewSceneProgram(p program.Program, c program.Config) *SceneProgram {
	return &SceneProgram{
		programRunner: p,
		programConfig: c,
	}
}

func (sp *SceneProgram) SetProgram(p program.Program) {
	sp.programRunner = p
}

func (sp *SceneProgram) SetProgramConfig(c program.Config) {
	sp.programConfig = c
}
