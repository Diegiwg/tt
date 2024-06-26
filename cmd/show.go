package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
)

func Show(ctx *cli.Context) error {
	table := data.ReadOrCreateRecord(ctx)

	r := table.GetLast()

	if len(r.Items) == 0 {
		return errors.New("no items found in record")
	}

	var acc time.Duration

	count := len(r.Items)
	for index, item := range r.Items {
		start := item.Start
		end := item.End

		if end == (time.Time{}) {

			if index != count-1 {
				continue
			}

			end = time.Now()
		}

		fmt.Printf("Item: %02d:%02d - %02d:%02d\n", start.Hour(), start.Minute(), end.Hour(), end.Minute())

		acc += end.Sub(start)
	}

	println("Total time:", acc.String()+"\n")

	return nil
}
