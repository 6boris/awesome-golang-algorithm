package Solution

func Solution(nums []int, k int) int {
	l := len(nums)
	minArr := make([]int, l)
	minArr[k] = nums[k]
	for i := k - 1; i >= 0; i-- {
		if nums[i] > minArr[i+1] {
			minArr[i] = minArr[i+1]
		} else {
			minArr[i] = nums[i]
		}
	}
	for j := k + 1; j < l; j++ {
		if nums[j] > minArr[j-1] {
			minArr[j] = minArr[j-1]
		} else {
			minArr[j] = nums[j]
		}
	}

	ans := -1
	for i := 0; i <= k; i++ {
		if i > 0 && minArr[i] == minArr[i-1] {
			continue
		}
		for j := l - 1; j >= k; j-- {
			if j < k-1 && minArr[j] == minArr[j+1] {
				continue
			}
			a := minArr[i]
			b := minArr[j]
			if b < a {
				a = b
			}
			if r := a * (j - i + 1); ans == -1 || r > ans {
				ans = r
			}
		}
	}
	return ans
}
