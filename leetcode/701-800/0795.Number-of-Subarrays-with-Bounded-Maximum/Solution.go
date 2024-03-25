package Solution

func Solution(nums []int, left int, right int) int {
	ans := 0
	// 1, 2, 2, 1 [2,3]
	//    2[1,2], [2]
	//       2 [2], [2,2], [1,2,2], | [2,2,1], [1,2,2,1], [2,1]
	//
	for i := range nums {
		if nums[i] >= left && nums[i] <= right {
			li := i
			for ; li >= 0 && nums[li] <= nums[i]; li-- {
			}
			left := i - li
			ans += left

			ri := i + 1
			for ; ri < len(nums) && nums[ri] < nums[i]; ri++ {
			}
			ri--
			if diff := ri - i; diff > 0 {
				ans += diff * left
			}
		}
	}
	return ans
}
