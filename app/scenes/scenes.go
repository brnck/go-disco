package scenes

import (
	"github.com/brnck/go-disco/app/output"
	"sync"
)

type Scene struct {
	sceneProgram []SceneProgram
}

func (s *Scene) Execute(o output.Output) {
	var wg sync.WaitGroup

	for _, program := range s.sceneProgram {
		wg.Add(1)
		go program.programRunner.Run(o, program.start, program.end, &wg)
	}

	wg.Wait()
	s.ClearProgramsList()
}

func (s *Scene) ClearProgramsList() {
	s.sceneProgram = s.sceneProgram[:0]
}

func (s *Scene) AddProgram(p *SceneProgram) {
	s.sceneProgram = append(s.sceneProgram, *p)
}

func New() *Scene {
	return &Scene{
		sceneProgram: make([]SceneProgram, 0),
	}
}
