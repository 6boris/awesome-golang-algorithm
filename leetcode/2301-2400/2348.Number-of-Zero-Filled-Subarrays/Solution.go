package Solution

func Solution(nums []int) int64 {
	ans := int64(0)

	zero := int64(0)
	for _, n := range nums {
		if n == 0 {
			zero++
			ans += zero
			continue
		}
		zero = 0
	}
	return ans
}
