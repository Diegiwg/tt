package data

import (
	"encoding/json"
	"os"

	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/model"
)

func ReadOrCreateRecord(ctx *cli.Context) model.RecordTable {
	dbPath := Path()

	_, err := os.Stat(dbPath)
	if err != nil {
		table := model.NewRecordTable()
		table.Add(model.NewRecord())

		SaveRecordToFile(ctx, &table)
		return table
	}

	fileContent, err := os.ReadFile(dbPath)
	if err != nil {
		panic(err)
	}

	var table model.RecordTable

	err = json.Unmarshal(fileContent, &table)
	if err != nil {
		println("Database file is corrupted. Creating new one.")

		table = model.NewRecordTable()
		table.Add(model.NewRecord())

		SaveRecordToFile(ctx, &table)
	}

	return table
}
