package Solution

import "sort"

func sortedSquares_1(nums []int) []int {
	ans := make([]int, 0)
	for _, v := range nums {
		ans = append(ans, v*v)
	}
	sort.Ints(ans)
	return ans
}

func sortedSquares_2(nums []int) []int {
	left, right, idx := 0, len(nums)-1, len(nums)-1
	ans := make([]int, len(nums))
	for left <= right && idx >= 0 {
		if l, r := nums[left]*nums[left], nums[right]*nums[right]; l > r {
			ans[idx] = l
			left++
		} else {
			ans[idx] = r
			right--
		}
		idx--
	}
	return ans
}
