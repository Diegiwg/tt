package model

import (
	"errors"
	"fmt"
	"time"
)

type RecordItem struct {
	Start time.Time
	End   time.Time
}

func NewRecordItem() RecordItem {
	return RecordItem{
		Start: time.Now(),
	}
}

type Record struct {
	Items       []RecordItem
	CurrentItem int
}

func NewRecord() Record {
	return Record{
		Items:       []RecordItem{},
		CurrentItem: -1,
	}
}

func (record *Record) Start() error {
	if record.CurrentItem != -1 {
		return errors.New("already have a current item")
	}

	item := NewRecordItem()
	record.Items = append(record.Items, item)
	record.CurrentItem = len(record.Items) - 1

	return nil
}

func (record *Record) Stop() error {
	if record.CurrentItem == -1 {
		return errors.New("no current item to stop")
	}

	record.Items[record.CurrentItem].End = time.Now()
	record.CurrentItem = -1

	return nil
}

func (record *Record) TotalTime() string {
	var total time.Duration

	item_count := len(record.Items)
	for index, item := range record.Items {

		if item.End == (time.Time{}) {
			if index != item_count-1 {
				continue
			}

			item.End = time.Now()
		}

		fmt.Printf("Item: %02d:%02d - %02d:%02d\n", item.Start.Hour(), item.Start.Minute(), item.End.Hour(), item.End.Minute())

		total += item.End.Sub(item.Start)
	}

	return total.String() + "\n"
}
