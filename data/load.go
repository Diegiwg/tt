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
		println("record not found, creating new one")
		r := model.NewRecordTable()
		SaveRecordToFile(ctx, &r)
		return r
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
