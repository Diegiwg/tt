package cmd

import (
	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
	"github.com/Diegiwg/tt/model"
)

func Start(ctx *cli.Context) error {
	r := model.NewRecord()

	err := r.Start()
	if err != nil {
		return err
	}

	table := data.ReadOrCreateRecord(ctx)
	table.Add(r)

	data.SaveRecordToFile(ctx, &table)

	return nil
}
