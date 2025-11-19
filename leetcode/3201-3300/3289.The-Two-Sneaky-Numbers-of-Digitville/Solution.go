package Solution

func Solution(nums []int) []int {
	ret := make([]int, 2)
	index := 0
	exists := make([]bool, len(nums)-2)
	for _, n := range nums {
		if exists[n] {
			ret[index] = n
			index++
			continue
		}
		exists[n] = true
	}
	return ret
}
