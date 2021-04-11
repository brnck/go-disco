package scenes

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/programs"
)

type SceneProgram struct {
	programRunner programs.Program
	programConfig config.Program
}

func (sp *SceneProgram) SetProgram(p programs.Program) {
	sp.programRunner = p
}

func (sp *SceneProgram) SetProgramConfig(c config.Program) {
	sp.programConfig = c
}

func NewSceneProgram() *SceneProgram {
	return &SceneProgram{}
}
