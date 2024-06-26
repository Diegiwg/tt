package cmd

import (
	"strconv"
	"time"

	"github.com/Diegiwg/cli"
	"github.com/Diegiwg/tt/data"
	"github.com/Diegiwg/tt/model"
)

func totalTime(r model.Record) string {
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

		acc += end.Sub(start)
	}

	return "Total time: " + acc.String()
}

func List(ctx *cli.Context) error {
	table := data.ReadOrCreateRecord(ctx)
	recordsInTable := len(table.Records)

	limit, err := strconv.Atoi(ctx.Flags["limit"])
	if err != nil {
		limit = 10
	}

	println("List of last " + strconv.Itoa(limit) + " records total time:")

	for index := recordsInTable - 1; index >= recordsInTable-limit; index-- {
		if index < 0 {
			break
		}

		r := table.Records[index]
		println(totalTime(r))
	}

	return nil
}
