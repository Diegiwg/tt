package data

import (
	"encoding/json"
	"os"

	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/model"
)

func SaveRecordToFile(ctx *cli.Context, table *model.RecordTable) {
	dbPath := Path()

	if len(table.Records) > DATA_BASE_RECORDS_LIMIT {
		table.Records = table.Records[len(table.Records)-DATA_BASE_RECORDS_LIMIT:]
	}

	data, err := json.Marshal(table)
	if err != nil {
		panic(err)
	}

	os.WriteFile(dbPath, []byte(data), 0644)
}
