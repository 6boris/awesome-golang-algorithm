package Solution

import "sort"

func Solution(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	ans := 0
	numCount := make(map[int]int)
	modes := make(map[int]bool)

	addMode := func(value int) {
		modes[value] = true
		if value-k >= nums[0] {
			modes[value-k] = true
		}
		if value+k <= nums[len(nums)-1] {
			modes[value+k] = true
		}
	}

	lastNumIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[lastNumIndex] {
			numCount[nums[lastNumIndex]] = i - lastNumIndex
			if i-lastNumIndex > ans {
				ans = i - lastNumIndex
			}
			addMode(nums[lastNumIndex])
			lastNumIndex = i
		}
	}

	numCount[nums[lastNumIndex]] = len(nums) - lastNumIndex
	if len(nums)-lastNumIndex > ans {
		ans = len(nums) - lastNumIndex
	}
	addMode(nums[lastNumIndex])

	leftBound := func(value int) int {
		left, right := 0, len(nums)-1
		for left < right {
			mid := (left + right) / 2
			if nums[mid] < value {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return left
	}

	rightBound := func(value int) int {
		left, right := 0, len(nums)-1
		for left < right {
			mid := (left + right + 1) / 2
			if nums[mid] > value {
				right = mid - 1
			} else {
				left = mid
			}
		}
		return left
	}

	uniqueModes := make([]int, 0, len(modes))
	for mode := range modes {
		uniqueModes = append(uniqueModes, mode)
	}
	sort.Ints(uniqueModes)

	for _, mode := range uniqueModes {
		l := leftBound(mode - k)
		r := rightBound(mode + k)
		var tempAns int
		if count, exists := numCount[mode]; exists {
			tempAns = min(r-l+1, count+numOperations)
		} else {
			tempAns = min(r-l+1, numOperations)
		}
		if tempAns > ans {
			ans = tempAns
		}
	}

	return ans
}
