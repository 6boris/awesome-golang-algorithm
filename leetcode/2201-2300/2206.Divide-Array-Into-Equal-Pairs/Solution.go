package Solution

func Solution(nums []int) bool {
	bucket := make([]int, 501)
	for _, n := range nums {
		bucket[n]++
	}
	for idx := 0; idx < 501; idx++ {
		if bucket[idx]&1 == 1 {
			return false
		}
	}
	return true
}
