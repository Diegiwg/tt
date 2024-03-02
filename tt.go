package main

import cli "github.com/Diegiwg/cli"

// tt start -> Cria um novo record na DB, salvando o tempo que o comando foi executado
// tt pause -> Adiciona um tempo de pausa no record
// tt resume -> Adiciona um marcador de tempo para voltar a contar
// tt stop -> Finaliza o record
// tt msg <msg> -> Adiciona uma mensagem ao record atual
// tt -> Mostra o status atual
// tt list -> Mostra todos os records
// tt show <id> -> Mostra um record

func main() {
	app := cli.NewApp()

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
