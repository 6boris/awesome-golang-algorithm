package Solution

import (
	"math"
)

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < len(nums); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 0; i < n-1; i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < n {
				dp[i+j] = Min(dp[i+j], dp[i]+1)
			}
		}
	}
	return dp[n-1]
}

func jump2(nums []int) int {
	//特殊情况判断
	if len(nums)==1{
		return 0
	}else if nums[0] >= len(nums){
		return 1
	}
	left, right, res := 0, nums[0], 0
	for right < len(nums)-1 {
		max := 0
		for i := left; i <= right; i++ {
			//寻找最大跨度
			if nums[i]-(right-i) >= max{
				max = nums[i]-(right - i)
			}
		}
		//窗口滑动右滑
		left = right
		right += max
		res++
	}
	//判断最后一步是不是踏在最后一个格子。循环条件结束只能确定能到达边界，left指针才是每次跳的格子。
	if left<len(nums)-1{
		res++
	}
	return res
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
