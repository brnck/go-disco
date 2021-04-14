package scenes

import (
	"github.com/brnck/go-disco/app/output"
	"sync"
)

type Scene struct {
	sceneProgram []SceneProgram
	interval     int
}

func (s *Scene) Execute(o output.Output) {
	var wg sync.WaitGroup

	for i := 0; i < s.interval; i++ {
		for _, program := range s.sceneProgram {
			wg.Add(1)
			go program.programRunner.Run(o, program.programConfig, &wg)
		}

		wg.Wait()
	}

	s.ClearProgramsList()
}

func (s *Scene) ClearProgramsList() {
	s.sceneProgram = s.sceneProgram[:0]
}

func (s *Scene) AddProgram(p *SceneProgram) {
	s.sceneProgram = append(s.sceneProgram, *p)
}

func (s *Scene) SetInterval(i int) {
	s.interval = i
}

func New() *Scene {
	return &Scene{
		sceneProgram: make([]SceneProgram, 0),
		interval:     1,
	}
}
