package app

import (
	"github.com/brnck/go-disco/app/output"
	"github.com/brnck/go-disco/app/output/ledstrip"
	"github.com/brnck/go-disco/app/output/stdout"
	"github.com/brnck/go-disco/app/utils"
	"log"
	"os"
)

type App struct {
	Output output.Output
	Logger *log.Logger
}

func (a *App) SetOutput(o output.Output) {
	a.Output = o
}

func (a *App) SetLogger(l *log.Logger) {
	a.Logger = l
}

func Init() (*App, error) {
	app := &App{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}

	standardOutput, err := initializeStandardOutput(app.Logger)

	if err != nil {
		return nil, err
	}
	app.SetOutput(standardOutput)

	return app, nil
}

func initializeStandardOutput(l *log.Logger) (*stdout.Stdout, error) {
	ws, err := stdout.Init(l)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func initializeWS2812Output() (*ledstrip.Led, error) {
	ws, err := ledstrip.Init()
	if err != nil {
		return nil, err
	}

	return ws, err
}

func (a *App) Run() {
	color := utils.RgbToColor(255, 255, 255)

	a.Output.SetLed(0, color)
}
