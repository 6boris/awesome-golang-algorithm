package Solution

func Solution(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		// l = 1
		ans += nums[i]
	}

	var selector func(int, int, int)
	selector = func(index, l, cur int) {
		if l == 0 {
			ans += cur
			return
		}
		if index == len(nums) {
			return
		}

		selector(index+1, l, cur)
		selector(index+1, l-1, cur^nums[index])
	}
	for l := 2; l <= len(nums); l++ {
		selector(0, l, 0)
	}
	return ans
}
