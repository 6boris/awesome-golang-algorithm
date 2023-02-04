package Solution

func Solution(nums []int, operations [][]int) []int {
	indies := make(map[int]int)
	for i, n := range nums {
		indies[n] = i
	}
	for _, op := range operations {
		i := indies[op[0]]
		nums[i] = op[1]
		indies[nums[i]] = i
	}
	return nums
}
