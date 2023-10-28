package Solution

func Solution(nums []int) []int {
	ans := make([]int, 0)
	for _, n := range nums {
		if n < 0 {
			n = -n
		}
		targetIndex := n - 1
		if nums[targetIndex] < 0 {
			ans = append(ans, n)
			continue
		}
		nums[targetIndex] = -nums[targetIndex]
	}
	return ans
}
