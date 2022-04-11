package Solution

func Solution(A []int) []bool {
	n := len(A)
	if n == 1 {
		if A[0] == 0 {
			return []bool{true}
		} else {
			return []bool{false}
		}
	}
	ans, v := make([]bool, n), 0
	for i := 0; i < n; i++ {
		v = (2*v + A[i]) % 5
		if v == 0 {
			ans[i] = true
		} else {
			ans[i] = false
		}
	}
	return ans
}

func Solution2(nums []int) []bool {
	base := make([]int, len(nums))
	base[0] = nums[0]
	result := make([]bool, len(nums))
	result[0] = nums[0] == 0
	for idx := 1; idx < len(nums); idx++ {
		base[idx] = nums[idx] + nums[idx-1]*2
		if idx >= 4 {
			base[idx] += base[idx-4]
		}
		cmp := base[idx]
		if idx >= 2 {
			cmp -= base[idx-2]
			if cmp < 0 {
				cmp = -cmp
			}
		}
		if cmp%5 == 0 {
			result[idx] = true
		}
	}
	return result
}
