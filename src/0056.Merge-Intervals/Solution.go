package Solution

import "sort"

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {

	sort.Sort()

	return []Interval{}

}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Sort(intervals []Interval) []Interval {
	i, j, tmp := 0, 0, Interval{Start: 0, End: 0}

	for i := 1; i < len(intervals); i++ {
		tmp = intervals[i]
		for j = i; j > 0 && Less(intervals[j],tmp); j--{
			intervals =
		}
	}
}

func Less(i, j Interval) bool {
	return false
}
