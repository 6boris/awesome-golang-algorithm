package Solution

func Solution(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	start := 0
	for nums[start] != 0 {
		start++
	}

	asc := true
	desc := true

	for i := 1; i < n; i++ {
		value := nums[(start+i)%n]

		if value != i {
			asc = false
		}
		if value != n-i {
			desc = false
		}

		if !asc && !desc {
			return -1
		}
	}

	if asc {
		return min(start, n-start+2)
	}

	return min(n-start, start+2)
}
