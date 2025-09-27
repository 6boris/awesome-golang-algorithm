package Solution

func Solution(logs [][]int, k int) []int {
	ret := make([]int, k)
	// 1:1
	// 2:2
	count := make(map[int]map[int]struct{})
	for _, log := range logs {
		if _, ok := count[log[0]]; !ok {
			count[log[0]] = make(map[int]struct{})
		}
		count[log[0]][log[1]] = struct{}{}
	}
	for _, c := range count {
		ret[len(c)-1]++
	}
	return ret
}
