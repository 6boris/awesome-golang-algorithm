package Solution

import (
	"sort"
)

//	dp[i] = dp[i-1] + 1

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp, ans := make([]int, len(nums)), 1

	//	子序列最少要包含自己
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(dp[i], ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLIS1(nums []int) int {
	tails := make([]int, len(nums)) // 长度为i的子数组最小的末尾数组
	var result int
	for x := 0; x < len(nums); x++ {
		var i, j int = 0, result
		for i != j {
			m := (i + j) / 2
			if tails[m] < nums[x] {
				i = m + 1
			} else {
				j = m
			}
		}
		tails[i] = nums[x]
		if i == result {
			result++
		}
	}
	return result
}

func lengthOfLIS2(nums []int) int {
	t := []int{}
	for _, n := range nums {
		j := sort.Search(len(t), func(i int) bool {
			return t[i] >= n
		})
		if j == len(t) {
			t = append(t, n)
		} else {
			t[j] = n
		}
	}
	return len(t)
}
