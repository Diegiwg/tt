package cmd

import (
	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
)

func Resume(ctx *cli.Context) error {
	table := data.ReadOrCreateRecord(ctx)

	record := table.GetLast()
	err := record.Start()
	if err != nil {
		return err
	}

	data.SaveRecordToFile(ctx, &table)

	return nil
}
