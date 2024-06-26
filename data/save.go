package data

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/model"
)

func SaveRecordToFile(ctx *cli.Context, record *model.Record) {
	USER_HOME, _ := os.UserHomeDir()

	dbPath := filepath.Join(USER_HOME, "tt.db")

	data, err := json.Marshal(record)
	if err != nil {
		panic(err)
	}

	os.WriteFile(dbPath, []byte(data), 0644)
}
