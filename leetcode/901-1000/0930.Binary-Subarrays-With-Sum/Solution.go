package Solution

func Solution(nums []int, goal int) int {
	count := map[int]int{0: 1}
	ans, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		diff := goal - sum
		if diff > 0 {
			count[sum]++
			continue
		}
		diff = -diff
		if r := count[diff]; r > 0 {
			ans += r
		}
		count[sum]++
	}

	return ans
}
