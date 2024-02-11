package Solution

func Solution(nums []int) int {
	sum := 0
	indies := make(map[int]int)
	ans := 0
	// 0 0 1 0 0 1 1

	// 0 1 0
	// -1:0
	//
	indies[0] = -1
	for i := 0; i < len(nums); i++ {
		add := 1
		if nums[i] == 0 {
			add = -1
		}
		sum += add
		if idx, ok := indies[sum]; ok {
			if r := i - idx; r > ans {
				ans = r
			}
			continue
		}
		indies[sum] = i
	}
	return ans
}
