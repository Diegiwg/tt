package data

import (
	"os"
	"path/filepath"
	"time"

	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/model"
)

func SaveRecordToFile(ctx *cli.Context, record *model.Record) {
	USER_HOME, _ := os.UserHomeDir()
	DB_PATH := filepath.Join(USER_HOME, "tt.db")

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

	os.WriteFile(DB_PATH, []byte(fileContent), 0644)
}
