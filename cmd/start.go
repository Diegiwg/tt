package cmd

import (
	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
	"github.com/Diegiwg/tt/model"
)

func Start(ctx *cli.Context) error {
	r := model.NewRecord()
	r.Start()

	data.SaveRecordToFile(ctx, &r)

	return nil
}
