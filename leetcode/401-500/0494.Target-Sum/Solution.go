package Solution

func Solution(nums []int, target int) int {
	ans := 0
	var dfs func(int, int)
	dfs = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				ans++
			}
			return
		}
		dfs(index+1, sum-nums[index])
		dfs(index+1, sum+nums[index])
	}
	dfs(0, 0)
	return ans
}
