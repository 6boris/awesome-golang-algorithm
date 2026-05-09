package Solution

func Solution(nums []int) int {
	maxLen := 0
	n := len(nums)

	// 外层循环：确定子数组的起点
	for left := 0; left < n; left++ {
		// 使用 map 来存储窗口内已经出现的不同数字
		evenSeen := make(map[int]struct{})
		oddSeen := make(map[int]struct{})

		// 内层循环：不断向右扩展窗口
		for right := left; right < n; right++ {
			num := nums[right]

			// 分类记录去重后的数字
			if num%2 == 0 {
				evenSeen[num] = struct{}{}
			} else {
				oddSeen[num] = struct{}{}
			}

			// 检查当前子数组是否平衡
			if len(evenSeen) == len(oddSeen) {
				currentLen := right - left + 1
				if currentLen > maxLen {
					maxLen = currentLen
				}
			}
		}
	}

	return maxLen
}
