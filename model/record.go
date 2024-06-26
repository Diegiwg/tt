package model

import (
	"errors"
	"fmt"
	"time"
)

type RecordItem struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type Record struct {
	Items       []RecordItem `json:"items"`
	CurrentItem int          `json:"current_item"`
}

type RecordTable struct {
	Records []Record `json:"records"`
}

func NewRecordItem() RecordItem {
	return RecordItem{
		Start: time.Now(),
	}
}

func NewRecord() Record {
	return Record{
		Items:       []RecordItem{},
		CurrentItem: -1,
	}
}

func NewRecordTable() RecordTable {
	return RecordTable{
		Records: []Record{},
	}
}

func (table *RecordTable) Add(record Record) {
	table.Records = append(table.Records, record)
}

func (table *RecordTable) GetLast() Record {
	return table.Records[len(table.Records)-1]
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

	count := len(record.Items)
	for index, item := range record.Items {

		if item.End == (time.Time{}) {
			if index != count-1 {
				continue
			}

			item.End = time.Now()
		}

		fmt.Printf("Item: %02d:%02d - %02d:%02d\n", item.Start.Hour(), item.Start.Minute(), item.End.Hour(), item.End.Minute())

		total += item.End.Sub(item.Start)
	}

	return total.String() + "\n"
}
