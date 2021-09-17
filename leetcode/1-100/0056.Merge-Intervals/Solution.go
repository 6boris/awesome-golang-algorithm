package Solution

import (
	"sort"
)

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

//	自定义排序规则
type SortByInt []Interval

func (p SortByInt) Len() int {
	return len(p)
}

func (p SortByInt) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p SortByInt) Less(i, j int) bool {
	return p[i].Start < p[j].Start
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//
func merge(intervals []Interval) []Interval {
	if intervals == nil || len(intervals) <= 1 {
		return intervals
	}
	//	先将结构体排序
	sort.Sort(SortByInt(intervals))

	start := intervals[0].Start
	end := intervals[0].End

	ans := make([]Interval, 0)

	for _, v := range intervals {
		if v.Start <= end {
			end = Max(end, v.End)
		} else {
			ans = append(ans, Interval{Start: start, End: end})
			start = v.Start
			end = v.End
		}
	}
	ans = append(ans, Interval{Start: start, End: end})
	return ans
}
