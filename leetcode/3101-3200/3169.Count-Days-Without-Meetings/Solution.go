package Solution

import "sort"

func Solution(days int, meetings [][]int) int {
	ans := 0
	sort.Slice(meetings, func(i, j int) bool {
		a, b := meetings[i], meetings[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	})
	end := 0
	for i := 0; i < len(meetings); i++ {
		if meetings[i][0] > end {
			ans += meetings[i][0] - end - 1
		}
		end = max(end, meetings[i][1])
	}
	ans += days - end
	return ans
}
