package main

import cli "github.com/Diegiwg/cli"

func main() {
	app := cli.NewApp()

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
