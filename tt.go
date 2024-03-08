package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	cli "github.com/Diegiwg/cli"
)

// tt start -> Cria um novo record na DB, salvando o tempo que o comando foi executado
// tt pause -> Adiciona um tempo de pausa no record
// tt resume -> Adiciona um marcador de tempo para voltar a contar
// tt stop -> Finaliza o record
// tt msg <msg> -> Adiciona uma mensagem ao record atual
// tt status -> Mostra o status atual
// tt list -> Mostra todos os records
// tt show <id> -> Mostra um record

type RecordItem struct {
	Start time.Time
	End   time.Time
}

func NewRecordItem() RecordItem {
	return RecordItem{
		Start: time.Now(),
	}
}

type Record struct {
	Items       []RecordItem
	CurrentItem int
}

func NewRecord() Record {
	return Record{
		Items:       []RecordItem{},
		CurrentItem: -1,
	}
}

func (record *Record) Start() error {
	if record.CurrentItem != -1 {
		return errors.New("already have a current item")
	}

	item := NewRecordItem()
	record.Items = append(record.Items, item)
	record.CurrentItem = len(record.Items) - 1

	return nil
}

func (record *Record) Stop() error {
	if record.CurrentItem == -1 {
		return errors.New("no current item to stop")
	}

	record.Items[record.CurrentItem].End = time.Now()
	record.CurrentItem = -1

	return nil
}

func (record *Record) TotalTime() string {
	var total time.Duration

	for _, item := range record.Items {

		if item.End == (time.Time{}) {
			continue
		}

		fmt.Printf("Item: %02d:%02d - %02d:%02d\n", item.Start.Hour(), item.Start.Minute(), item.End.Hour(), item.End.Minute())

		total += item.End.Sub(item.Start)
	}

	return total.String() + "\n"
}

func SaveRecordToFile(ctx *cli.Context, record *Record) {
	dbPath := filepath.Join(filepath.Dir(ctx.App.Program), "tt.db")

	fileContent := ""

	// Convert All Items to String
	for _, item := range record.Items {
		// fileContent += item.Start.Format(time.RFC3339) + " " + item.End.Format(time.RFC3339) + "\n"
		fileContent += item.Start.Format(time.RFC3339) + "$"

		if item.End == (time.Time{}) {
			fileContent += "\n"
		} else {
			fileContent += item.End.Format(time.RFC3339) + "\n"
		}
	}

	os.WriteFile(dbPath, []byte(fileContent), 0644)
}

func ReadOrCreateRecord(ctx *cli.Context) Record {

	dbPath := filepath.Join(filepath.Dir(ctx.App.Program), "tt.db")

	_, err := os.Stat(dbPath)
	if err != nil {
		println("record not found, creating new one")
		r := NewRecord()
		SaveRecordToFile(ctx, &r)
		return r
	}

	fileContent, err := os.ReadFile(dbPath)
	if err != nil {
		panic(err)
	}

	record := NewRecord()
	for _, line := range strings.Split(string(fileContent), "\n") {
		if line == "" {
			continue
		}

		splitValues := strings.Split(line, "$")

		start, _ := time.Parse(time.RFC3339, splitValues[0])

		end, err := time.Parse(time.RFC3339, splitValues[1])
		if err != nil {
			end = time.Time{}
		}

		record.Items = append(record.Items, RecordItem{
			Start: start,
			End:   end,
		})
	}

	// Check if exist a current item
	for i, item := range record.Items {
		if item.End == (time.Time{}) {
			record.CurrentItem = i
			break
		}
	}

	return record
}

func cmdStart(ctx *cli.Context) error {
	r := NewRecord()
	r.Start()

	SaveRecordToFile(ctx, &r)

	return nil
}

func cmdResume(ctx *cli.Context) error {
	r := ReadOrCreateRecord(ctx)
	r.Start()
	SaveRecordToFile(ctx, &r)

	return nil
}

func cmdPause(ctx *cli.Context) error {
	r := ReadOrCreateRecord(ctx)
	r.Stop()
	SaveRecordToFile(ctx, &r)

	return nil
}

func cmdStop(ctx *cli.Context) error {
	r := ReadOrCreateRecord(ctx)

	if r.CurrentItem != -1 && r.Stop() != nil {
		return errors.New("no current item to stop")
	}

	time := r.TotalTime()
	println("Total time:", time)

	r = NewRecord()
	SaveRecordToFile(ctx, &r)

	return nil
}

func main() {
	app := cli.NewApp()

	app.AddCommand(&cli.Command{
		Name:  "start",
		Desc:  "Start a record of time",
		Help:  "Start a record of time",
		Usage: "",
		Exec:  cmdStart,
	})

	app.AddCommand(&cli.Command{
		Name:  "pause",
		Desc:  "Pause a record of time",
		Help:  "Pause a record of time",
		Usage: "",
		Exec:  cmdPause,
	})

	app.AddCommand(&cli.Command{
		Name:  "resume",
		Desc:  "Resume a record of time",
		Help:  "Resume a record of time",
		Usage: "",
		Exec:  cmdResume,
	})

	app.AddCommand(&cli.Command{
		Name:  "stop",
		Desc:  "Stop a record of time",
		Help:  "Stop a record of time",
		Usage: "",
		Exec:  cmdStop,
	})

	err := app.Run()
	if err != nil {
		println(err.Error())
	}
}
