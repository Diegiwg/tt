package cmd

import (
	"errors"

	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
	"github.com/Diegiwg/tt/model"
)

func Stop(ctx *cli.Context) error {
	r := data.ReadOrCreateRecord(ctx)

	if r.CurrentItem != -1 && r.Stop() != nil {
		return errors.New("no current item to stop")
	}

	time := r.TotalTime()
	println("Total time:", time)

	r = model.NewRecord()
	data.SaveRecordToFile(ctx, &r)

	return nil
}
