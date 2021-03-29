package main

import (
	application "github.com/brnck/go-disco/app"
)

func main() {
	app, err := application.Init()
	if err != nil {
		panic(err)
	}
	if err := app.Output.Open(); err != nil {
		panic(err)
	}
	defer app.Output.Close()

	app.Run()
}
