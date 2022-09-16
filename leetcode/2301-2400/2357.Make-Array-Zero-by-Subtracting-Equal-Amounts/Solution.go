package Solution

func Solution(nums []int) int {
	bucket := make(map[int]struct{})
	for _, n := range nums {
		if n == 0 {
			continue
		}
		bucket[n] = struct{}{}
	}
	return len(bucket)
}
