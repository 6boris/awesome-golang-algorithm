package Solution

func Solution(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	// 用下表跳?
	for idx := 0; idx < len(temperatures)-1; idx++ {
		next := idx + 1
		for ; next < len(temperatures) && temperatures[next] <= temperatures[idx]; next++ {
		}
		if next < len(temperatures) {
			ans[idx] = next - idx
		}
	}

	return ans
}
