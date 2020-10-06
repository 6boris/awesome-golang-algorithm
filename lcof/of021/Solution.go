package Solution

func exchange(nums []int) []int {
	start, end := 0, len(nums)-1

	for start < end {
		for start < end && (nums[start]%2 == 1) {
			start++
		}
		for start < end && (nums[end]%2 == 0) {
			end--
		}
		nums[start], nums[end] = nums[end], nums[start]
	}
	return nums
}

func exchange2(nums []int) []int {
	ans := make([]int, 0, 0)
	for _, v := range nums {
		if v&1 != 0 {
			ans = append(ans, v)
		}
	}
	for _, v := range nums {
		if v&1 == 0 {
			ans = append(ans, v)
		}
	}
	return ans
}

func exchange3(nums []int) []int {
	i := 0

	for j := 0; j < len(nums); j++ {
		if nums[j]&1 == 1 {
			tmp := nums[j]
			for k := j - 1; k >= i; k-- {
				nums[k+1] = nums[k]
			}
			nums[i] = tmp
			i += 1
		}
	}
	return nums
}
