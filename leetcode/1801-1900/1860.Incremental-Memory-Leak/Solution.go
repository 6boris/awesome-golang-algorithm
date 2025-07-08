package Solution

func Solution(memory1 int, memory2 int) []int {
	ans := []int{1, memory1, memory2}
	for ans[0] <= ans[1] || ans[0] <= ans[2] {
		if ans[1] >= ans[2] {
			ans[1] -= ans[0]
		} else {
			ans[2] -= ans[0]
		}
		ans[0]++
	}
	return ans
}
