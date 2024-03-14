package main

import (
	cli "github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/cmd"
)

// tt start -> Cria um novo record na DB, salvando o tempo que o comando foi executado
// tt pause -> Adiciona um tempo de pausa no record
// tt resume -> Adiciona um marcador de tempo para voltar a contar
// tt stop -> Finaliza o record
// tt msg <msg> -> Adiciona uma mensagem ao record atual
// tt status -> Mostra o status atual
// tt list -> Mostra todos os records
// tt show <id> -> Mostra um record

func main() {
	app := cli.NewApp()

	app.AddCommand(&cli.Command{
		Name:  "start",
		Desc:  "Start a record of time",
		Help:  "Start a record of time",
		Usage: "",
		Exec:  cmd.Start,
	})

	app.AddCommand(&cli.Command{
		Name:  "pause",
		Desc:  "Pause a record of time",
		Help:  "Pause a record of time",
		Usage: "",
		Exec:  cmd.Pause,
	})

	app.AddCommand(&cli.Command{
		Name:  "resume",
		Desc:  "Resume a record of time",
		Help:  "Resume a record of time",
		Usage: "",
		Exec:  cmd.Resume,
	})

	app.AddCommand(&cli.Command{
		Name:  "stop",
		Desc:  "Stop a record of time",
		Help:  "Stop a record of time",
		Usage: "",
		Exec:  cmd.Stop,
	})

	app.AddCommand(&cli.Command{
		Name:  "show",
		Desc:  "Show a record of time",
		Help:  "Show a record of time",
		Usage: "",
		Exec:  cmd.Show,
	})

	err := app.Run()
	if err != nil {
		println(err.Error())
	}
}
