package main

import (
	application "github.com/brnck/go-disco/app"
	"os"
	"os/signal"
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			app.Logger.Printf("Caught signal %v. Exiting", sig)
			os.Exit(0)
		}
	}()

	for {
		app.Logger.Printf("Running iteration")
		if err := app.Run(); err != nil {
			panic(err)
		}
	}

}
