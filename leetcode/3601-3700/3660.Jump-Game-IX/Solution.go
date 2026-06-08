package Solution

import "math"

func Solution(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	// [value, index] 对
	type Item struct {
		value int
		index int
	}
	prevMax := make([]Item, n)

	prev := Item{math.MinInt32, -1}
	for i := 0; i < n; i++ {
		if nums[i] > prev.value {
			prev = Item{nums[i], i}
		}
		prevMax[i] = prev
	}

	var process func(r, rightMin, rightMax int)
	process = func(r, rightMin, rightMax int) {
		pMax := prevMax[r].value
		pivotIndex := prevMax[r].index

		var currMax int
		if pMax <= rightMin {
			currMax = pMax
		} else {
			currMax = rightMax
		}

		nextRightMin := rightMin
		if pMax < nextRightMin {
			nextRightMin = pMax
		}
		for i := pivotIndex; i <= r; i++ {
			ans[i] = currMax
			if nums[i] < nextRightMin {
				nextRightMin = nums[i]
			}
		}

		if pivotIndex == 0 {
			return
		}

		process(pivotIndex-1, nextRightMin, currMax)
	}

	process(n-1, math.MaxInt32, 0)

	return ans
}
