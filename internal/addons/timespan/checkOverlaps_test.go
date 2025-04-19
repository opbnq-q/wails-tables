package timespan

import (
	"testing"
	"time"
)

type Flight struct {
	Id        uint `gorm:"primaryKey"`
	StartTime int64
	EndTime   int64
}

var items = []*Flight{
	{Id: 1, StartTime: time.Now().Unix(), EndTime: time.Now().Add(time.Hour * 24).Unix()},
	{Id: 2, StartTime: time.Now().Add(time.Hour * 25).Unix(), EndTime: time.Now().Add(time.Hour * 28).Unix()},
	{Id: 3, StartTime: time.Now().Unix(), EndTime: time.Now().Add(time.Hour * 12).Unix()},
}

func TestCheckNonOverlapping(t *testing.T) {
	if CheckNonOverlapping(items, func(item *Flight) (int64, int64) {
		return item.StartTime, item.EndTime
	}) {
		t.Error("False positive on overlapping timespans")
	}

	nonOverlappingItems := items[:1]

	if !CheckNonOverlapping(nonOverlappingItems, func(item *Flight) (int64, int64) {
		return item.StartTime, item.EndTime
	}) {
		t.Error("False negative on non-overlapping timespans")
	}
}
