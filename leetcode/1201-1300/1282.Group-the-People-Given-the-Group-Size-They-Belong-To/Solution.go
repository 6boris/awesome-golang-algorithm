package Solution

func Solution(groupSizes []int) [][]int {
	groupBy := make(map[int][]int)
	for i, gs := range groupSizes {
		if _, ok := groupBy[gs]; !ok {
			groupBy[gs] = make([]int, 0)
		}
		groupBy[gs] = append(groupBy[gs], i)
	}
	ans := make([][]int, 0)
	for gs, people := range groupBy {
		for i := gs; i <= len(people); i += gs {
			ans = append(ans, people[i-gs:i])
		}
	}
	return ans
}
