package scenes

import "github.com/brnck/go-disco/app/programs"

type SceneProgram struct {
	start         int
	end           int
	programRunner programs.Program
}

func (sp *SceneProgram) SetStart(start int) {
	sp.start = start
}

func (sp *SceneProgram) SetEnd(end int) {
	sp.end = end
}

func (sp *SceneProgram) SetProgram(p programs.Program) {
	sp.programRunner = p
}

func NewSceneProgram() *SceneProgram {
	return &SceneProgram{}
}
