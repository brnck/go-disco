package app

import (
	"github.com/brnck/go-disco/app/config"
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/programs"
	"github.com/brnck/go-disco/app/scenes"
	"log"
	"os"
)

type App struct {
	Output   output.Output
	Config   *config.Config
	Logger   *log.Logger
	programs *programs.Programs
	scenes   *scenes.Scene
}

func (a *App) setOutput(o output.Output) {
	a.Output = o
}

func (a *App) setLogger(l *log.Logger) {
	a.Logger = l
}

func (a *App) setConfig(c *config.Config) {
	a.Config = c
}

func Init() (*App, error) {
	app := &App{
		Logger:   log.New(os.Stdout, "", log.LstdFlags),
		programs: programs.New(),
		scenes:   scenes.New(),
	}

	cfg, err := config.ParseConfig()
	if err != nil {
		return nil, err
	}
	app.setConfig(cfg)

	out, err := output.InitializeOutput(app.Config, app.Logger)
	if err != nil {
		return nil, err
	}
	app.setOutput(out)

	if err := registerPrograms(app); err != nil {
		return nil, err
	}

	return app, nil
}

func registerPrograms(app *App) error {
	if err := app.programs.AddProgram(programs.NewTheaterChase()); err != nil {
		return err
	}
	if err := app.programs.AddProgram(programs.NewTheaterRainbowChase()); err != nil {
		return err
	}
	if err := app.programs.AddProgram(programs.NewRangedRainbowCycle()); err != nil {
		return err
	}
	if err := app.programs.AddProgram(programs.NewStrobe()); err != nil {
		return err
	}
	if err := app.programs.AddProgram(programs.NewMapper()); err != nil {
		return err
	}
	if err := app.programs.AddProgram(programs.NewChaosColors()); err != nil {
		return err
	}
	if err := app.programs.AddProgram(programs.NewChaosFillDown()); err != nil {
		return err
	}
	if err := app.programs.AddProgram(programs.NewColorChase()); err != nil {
		return err
	}
	if err := app.programs.AddProgram(programs.NewColorChaseReverse()); err != nil {
		return err
	}
	if err := app.programs.AddProgram(programs.NewFadeRgb()); err != nil {
		return err
	}
	if err := app.programs.AddProgram(programs.NewFadeInOut()); err != nil {
		return err
	}

	return nil
}

func (a *App) Run() error {
	for _, scene := range a.Config.Scenes {
		sp := scenes.NewSceneProgram()
		for _, program := range scene.Programs {
			p, err := a.programs.GetProgram(program.Key)
			if err != nil {
				return nil
			}
			sp.SetProgram(p)
			sp.SetProgramConfig(program)
			a.scenes.AddProgram(sp)
		}
		a.scenes.Execute(a.Output)
	}

	return nil
}
