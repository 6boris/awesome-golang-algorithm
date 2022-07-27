package Solution

// top 2
func Solution(nums []int) int {
	ans := 0
	for idx := 0; idx < len(nums)-1; idx++ {
		for next := idx + 1; next < len(nums); next++ {
			if r := (nums[idx] - 1) * (nums[next] - 1); r > ans {
				ans = r
			}
		}
	}
	return ans
}
