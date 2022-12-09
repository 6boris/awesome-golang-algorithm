package Solution

// top 2
func Solution(nums []int, k int) bool {
	oneIdx := -1
	for i, n := range nums {
		if n == 1 {
			if oneIdx == -1 || i-oneIdx > k {
				// 满足条件
				oneIdx = i
				continue
			}
			return false
		}
	}
	return true
}
