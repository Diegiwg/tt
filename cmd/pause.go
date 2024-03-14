package cmd

import (
	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
)

func Pause(ctx *cli.Context) error {
	r := data.ReadOrCreateRecord(ctx)
	r.Stop()
	data.SaveRecordToFile(ctx, &r)

	return nil
}
