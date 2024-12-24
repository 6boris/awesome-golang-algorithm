package Solution

func Solution(nums []int) int {
	m := make(map[int]int)
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			cur := nums[i] * nums[j]
			m[cur]++
		}
	}
	ans := 0
	for _, n := range m {
		ans += (n - 1) * n / 2 * 8
	}
	return ans
}
