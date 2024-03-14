package main

import (
	cli "github.com/Diegiwg/cli"
	cmd "github.com/Diegiwg/tt/cmd"
)

func main() {
	app := cli.NewApp()

	app.AddCommand(&cli.Command{
		Name:  "start",
		Desc:  "Starts a new time record",
		Help:  "Clears the database and starts a new time record.",
		Usage: "",
		Exec:  cmd.Start,
	})

	app.AddCommand(&cli.Command{
		Name:  "pause",
		Desc:  "Adds a pause to the time record",
		Help:  "When you want to take a break, use this command.",
		Usage: "",
		Exec:  cmd.Pause,
	})

	app.AddCommand(&cli.Command{
		Name:  "resume",
		Desc:  "Resumes counting time after a pause",
		Help:  "When you want to resume after a pause, use this command.",
		Usage: "",
		Exec:  cmd.Resume,
	})

	app.AddCommand(&cli.Command{
		Name:  "stop",
		Desc:  "Stops time counting, clearing the time record",
		Help:  "When you want to finish the time record, use this command.",
		Usage: "",
		Exec:  cmd.Stop,
	})

	app.AddCommand(&cli.Command{
		Name:  "show",
		Desc:  "Shows the time passed in the time record",
		Help:  "When you want to see how much time has passed without finishing the record, use this command.",
		Usage: "",
		Exec:  cmd.Show,
	})

	err := app.Run()
	if err != nil {
		println(err.Error())
	}
}
