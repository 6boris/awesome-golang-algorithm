package Solution

func findContinuousSequence(target int) [][]int {
	left, right, ans := 1, 2, [][]int{}
	for left < right {
		sum := (left + right) * (right - left + 1) / 2
		if sum == target {
			arr := []int{}
			for i := left; i <= right; i++ {
				arr = append(arr, i)
			}
			ans = append(ans, arr)
			left++
		} else if sum < target {
			right++
		} else {
			left++
		}
	}
	return ans
}
