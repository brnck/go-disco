package app

import (
	"github.com/brnck/go-disco/pkg/output"
	"github.com/brnck/go-disco/pkg/program"
	scenes "github.com/brnck/go-disco/pkg/scene"
	"log"
	"os"
)

type App struct {
	Output   output.Output
	Config   *Config
	Logger   *log.Logger
	programs *program.Programs
	scenes   *scenes.Scene
}

func (a *App) setOutput(o output.Output) {
	a.Output = o
}

func (a *App) setLogger(l *log.Logger) {
	a.Logger = l
}

func (a *App) setConfig(c *Config) {
	a.Config = c
}

func (a *App) setPrograms(p *program.Programs) {
	a.programs = p
}

func New() (*App, error) {
	app := &App{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
		scenes: scenes.New(),
	}

	cfg, err := parseConfig()
	if err != nil {
		return nil, err
	}
	app.setConfig(cfg)

	out, err := output.InitializeOutput(app.Config.Output, buildOutputConfig(app.Config), app.Logger)
	if err != nil {
		return nil, err
	}
	app.setOutput(out)

	programs, err := program.New()
	if err != nil {
		return nil, err
	}
	app.setPrograms(programs)

	return app, nil
}

func (a *App) Run() error {
	for _, scene := range a.Config.Scenes {
		for _, liveProgram := range scene.Programs {
			p, err := a.programs.GetProgram(liveProgram.Key)
			if err != nil {
				return nil
			}
			sp := scenes.NewSceneProgram(p, liveProgram)
			a.scenes.AddProgram(sp)
			a.scenes.SetInterval(scene.Interval)
		}
		a.scenes.Execute(a.Output)
	}

	return nil
}
