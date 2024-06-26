package cmd

import (
	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
)

func Pause(ctx *cli.Context) error {
	table := data.ReadOrCreateRecord(ctx)

	r := table.GetLast()
	r.Stop()

	data.SaveRecordToFile(ctx, &table)

	return nil
}
