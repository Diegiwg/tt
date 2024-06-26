package cmd

import (
	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
)

func Pause(ctx *cli.Context) error {
	table := data.ReadOrCreateRecord(ctx)

	record := table.GetLast()

	err := record.Stop()
	if err != nil {
		return err
	}

	data.SaveRecordToFile(ctx, &table)

	return nil
}
