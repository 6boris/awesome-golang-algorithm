package Solution

func Solution(n int) int {
	used := make([]bool, n+1)
	var lookup func(int)
	ans := 0
	lookup = func(index int) {
		if index > n {
			ans++
			return
		}
		for idx := 1; idx <= n; idx++ {
			if !used[idx] && (index%idx == 0 || idx%index == 0) {
				used[idx] = true
				lookup(index + 1)
				used[idx] = false
			}
		}
	}
	lookup(1)
	return ans
}
