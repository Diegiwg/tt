package cmd

import (
	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
)

func Resume(ctx *cli.Context) error {
	r := data.ReadOrCreateRecord(ctx)
	r.Start()
	data.SaveRecordToFile(ctx, &r)

	return nil
}
