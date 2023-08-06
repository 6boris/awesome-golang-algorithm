package Solution

import (
	"sort"
)

type ScheduleData struct {
	startTime, endTime, profit []int
}

func (s *ScheduleData) Len() int {
	return len(s.profit)
}

func (s *ScheduleData) Swap(i, j int) {
	s.startTime[i], s.startTime[j] = s.startTime[j], s.startTime[i]
	s.endTime[i], s.endTime[j] = s.endTime[j], s.endTime[i]
	s.profit[i], s.profit[j] = s.profit[j], s.profit[i]
}

func (s *ScheduleData) Less(i, j int) bool {
	return s.endTime[i] < s.endTime[j]
}

func Solution(startTime []int, endTime []int, profit []int) int {
	l := len(profit)

	sd := ScheduleData{
		startTime: startTime,
		endTime:   endTime,
		profit:    profit,
	}
	sort.Sort(&sd)

	var bsearch func(int, int) int

	bsearch = func(end, target int) int {
		l, r := 0, end
		ans := -1
		for l < r {
			mid := (r-l)/2 + l
			if endTime[mid] <= startTime[end] {
				l = mid + 1
				ans = mid
				continue
			}
			r = mid
		}
		return ans
	}

	dp := make([]int, l)
	dp[0] = profit[0]
	ans := 0
	for idx := 1; idx < l; idx++ {
		index := bsearch(idx, startTime[idx])
		dp[idx] = profit[idx]

		if index < 0 || index == idx {
			if dp[idx] > ans {
				ans = dp[idx]
			}
			continue
		}
		for pre := index; pre >= 0; pre-- {
			if r := dp[pre] + profit[idx]; r > dp[idx] {
				dp[idx] = r
			}
		}
		if dp[idx] > ans {
			ans = dp[idx]
		}
	}
	return ans
}
