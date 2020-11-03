package Solution

//	动态规划
func lengthOfLIS(nums []int) int {
	dp, ans := []int{}, 0
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 1)
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

//	贪心 + 动态规划	
func lengthOfLIS2(nums []int) int {
	dp, ans := make([]int, len(nums)), 0
	for _, v := range nums {
		i := 0
		j := ans
		for i < j {
			m := (i + j) >> 1
			if dp[m] < v {
				i = m + 1
			} else {
				j = m
			}
		}
		dp[i] = v
		if ans == j {
			ans++
		}
	}

	return ans
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
