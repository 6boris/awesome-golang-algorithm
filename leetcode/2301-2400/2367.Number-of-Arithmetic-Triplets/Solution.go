package Solution

func Solution(nums []int, diff int) int {
	count := 0

	bucket := make([]int, 201)
	for i, n := range nums {
		bucket[n] = i + 1
	}

	for idx := 0; idx <= len(nums)-3; idx++ {
		next := nums[idx] + diff
		if !(next <= 200 && bucket[next] > 0) {
			continue
		}
		nextNext := next + diff
		if !(nextNext <= 200 && bucket[nextNext] > 0) {
			continue
		}
		count++
	}
	return count
}
