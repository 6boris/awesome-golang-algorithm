package Solution

func Solution(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	minIndex, maxIndex := 0, 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[minIndex] {
			minIndex = i
		}
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	left, right := minIndex, maxIndex
	if minIndex > maxIndex {
		left, right = maxIndex, minIndex
	}

	ans1 := right + 1
	ans2 := n - left
	ans3 := left + 1 + n - right
	return min(ans1, ans2, ans3)
}
