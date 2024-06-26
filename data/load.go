package data

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/model"
)

func ReadOrCreateRecord(ctx *cli.Context) model.RecordTable {
	USER_HOME, _ := os.UserHomeDir()

	dbPath := filepath.Join(USER_HOME, "tt.db")

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
		panic(err)
	}

	return table
}
