package data

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/model"
)

func ReadOrCreateRecord(ctx *cli.Context) model.Record {
	USER_HOME, _ := os.UserHomeDir()
	DB_PATH := filepath.Join(USER_HOME, "tt.db")

	_, err := os.Stat(DB_PATH)
	if err != nil {
		println("record not found, creating new one")
		r := model.NewRecord()
		SaveRecordToFile(ctx, &r)
		return r
	}

	fileContent, err := os.ReadFile(DB_PATH)
	if err != nil {
		panic(err)
	}

	record := model.NewRecord()
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

		record.Items = append(record.Items, model.RecordItem{
			Start: start,
			End:   end,
		})
	}

	// Check if exist a current item
	item_count := len(record.Items)
	for index, item := range record.Items {
		if item.End == (time.Time{}) {
			if index != item_count-1 {
				fmt.Println("End not found for Item", index+1)
				continue
			}

			record.CurrentItem = index
			break
		}
	}

	return record
}
