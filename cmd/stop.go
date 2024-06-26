package cmd

import (
	"errors"

	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
)

func Stop(ctx *cli.Context) error {
	table := data.ReadOrCreateRecord(ctx)

	r := table.GetLast()

	if r.CurrentItem != -1 && r.Stop() != nil {
		return errors.New("no current item to stop")
	}

	time := r.TotalTime()
	println("Total time:", time)

	data.SaveRecordToFile(ctx, &table)

	return nil
}
