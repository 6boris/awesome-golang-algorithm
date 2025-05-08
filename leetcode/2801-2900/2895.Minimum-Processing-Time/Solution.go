package Solution

import "sort"

func Solution(processorTime []int, tasks []int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] > tasks[j]
	})
	sort.Ints(processorTime)
	ans := 0
	index := 0
	for _, n := range processorTime {
		mm := max(tasks[index], tasks[index+1], tasks[index+2], tasks[index+3])
		index += 4
		ans = max(ans, n+mm)
	}
	return ans
}
