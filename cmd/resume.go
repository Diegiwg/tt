package cmd

import (
	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
)

func Resume(ctx *cli.Context) error {
	table := data.ReadOrCreateRecord(ctx)

	r := table.GetLast()
	r.Start()

	data.SaveRecordToFile(ctx, &table)

	return nil
}
