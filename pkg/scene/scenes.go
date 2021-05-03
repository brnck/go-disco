package scene

import (
	"github.com/brnck/go-disco/pkg/output"
	"github.com/brnck/go-disco/pkg/program"
	"sync"
)

type Config struct {
	Name     string           `mapstructure:"name"`
	Interval int              `mapstructure:"interval"`
	Programs []program.Config `mapstructure:"programs"`
}

type Scene struct {
	sceneProgram []SceneProgram
	interval     int
}

func New() *Scene {
	return &Scene{
		sceneProgram: make([]SceneProgram, 0),
		interval:     1,
	}
}

func (s *Scene) Execute(o output.Output) {
	var wg sync.WaitGroup

	for i := 0; i < s.interval; i++ {
		for _, liveProgram := range s.sceneProgram {
			wg.Add(1)
			go liveProgram.programRunner.Run(o, liveProgram.programConfig, &wg)
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
