package Solution

import (
	"math"
	"sort"
)

//	二分
func findDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

//	排序
func findDuplicate2(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return 0
}

//	map
func findDuplicate3(nums []int) int {
	found := -1
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > 1 {
			found = v
			break
		}
	}
	return found
}

//	排序
func findDuplicate4(nums []int) int {
	for i := range nums {
		tmp := int(math.Abs(float64(nums[i])))
		if nums[tmp-1] < 0 {
			return int(math.Abs(float64(nums[i])))
		}
		nums[tmp-1] = -nums[tmp-1]
	}
	return -1
}
