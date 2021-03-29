package app

import (
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/programs"
	"log"
	"os"
)

type App struct {
	Output   output.Output
	Logger   *log.Logger
	programs *programs.Programs
}

func (a *App) SetOutput(o output.Output) {
	a.Output = o
}

func (a *App) SetLogger(l *log.Logger) {
	a.Logger = l
}

func Init() (*App, error) {
	app := &App{
		Logger:   log.New(os.Stdout, "", log.LstdFlags),
		programs: programs.New(),
	}

	standardOutput, err := output.InitializeStandardOutput(app.Logger)

	if err != nil {
		return nil, err
	}
	app.SetOutput(standardOutput)

	if err := initializePrograms(app); err != nil {
		return nil, err
	}

	return app, nil
}

func initializePrograms(app *App) error {
	tc := programs.NewTheaterChase()
	trc := programs.NewTheaterRainbowChase()

	if err := app.programs.AddProgram(tc.Name, tc); err != nil {
		return err
	}
	if err := app.programs.AddProgram(trc.Name, trc); err != nil {
		return err
	}

	return nil
}

func (a *App) Run() {
	a.programs.RunProgram("theater_chase", a.Output)
	a.programs.RunProgram("theater_rainbow_chase", a.Output)
}
