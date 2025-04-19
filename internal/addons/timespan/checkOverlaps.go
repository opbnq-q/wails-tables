package timespan

import "sort"

func CheckNonOverlapping[T any](objects []*T, getInterval func(*T) (int64, int64)) bool {
	if len(objects) <= 1 {
		return true
	}

	intervals := make([]struct{ start, end int64 }, len(objects))
	for i, obj := range objects {
		start, end := getInterval(obj)
		if start > end {
			return false
		}
		intervals[i] = struct{ start, end int64 }{start, end}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i].start < intervals[i-1].end {
			return false
		}
	}

	return true
}
