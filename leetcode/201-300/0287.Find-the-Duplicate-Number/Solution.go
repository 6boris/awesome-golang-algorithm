package Solution

// 二分
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

func findDuplicate5(nums []int) int {
	ans := -1
	for idx := range nums {
		x := nums[idx]
		if x < 0 {
			x = -x
		}
		index := x - 1
		r := nums[index]
		if r < 0 {
			ans = x
			break
		}

		nums[index] = -nums[index]
	}
	for idx := range nums {
		if nums[idx] < 0 {
			nums[idx] = -nums[idx]
		}
	}
	return ans
}
