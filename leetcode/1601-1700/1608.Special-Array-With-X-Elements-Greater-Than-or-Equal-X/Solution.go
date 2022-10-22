package Solution

func Solution(nums []int) int {
	bucket := make([]int, 1001)
	for _, n := range nums {
		bucket[n]++
	}
	for idx := 999; idx >= 0; idx-- {
		bucket[idx] += bucket[idx+1]
		if bucket[idx] == idx {
			return idx
		}
	}
	return -1
}
