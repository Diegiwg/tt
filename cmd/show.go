package cmd

import (
	"errors"

	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
)

func Show(ctx *cli.Context) error {
	r := data.ReadOrCreateRecord(ctx)

	if len(r.Items) == 0 {
		return errors.New("no items found in record")
	}

	time := r.TotalTime()
	println("Total time:", time)

	return nil
}
