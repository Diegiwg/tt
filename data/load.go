package data

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/model"
)

func ReadOrCreateRecord(ctx *cli.Context) model.Record {
	USER_HOME, _ := os.UserHomeDir()

	dbPath := filepath.Join(USER_HOME, "tt.db")

	_, err := os.Stat(dbPath)
	if err != nil {
		println("record not found, creating new one")
		r := model.NewRecord()
		SaveRecordToFile(ctx, &r)
		return r
	}

	fileContent, err := os.ReadFile(dbPath)
	if err != nil {
		panic(err)
	}

	var record model.Record

	err = json.Unmarshal(fileContent, &record)
	if err != nil {
		panic(err)
	}

	return record
}
