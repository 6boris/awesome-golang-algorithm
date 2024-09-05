package Solution

func Solution(rolls []int, mean int, n int) []int {
	target := (n + len(rolls)) * mean
	for _, r := range rolls {
		target -= r
	}

	if n*6 < target || target < n {
		return []int{}
	}
	ans := make([]int, n)
	base := target / n
	for i := range n {
		ans[i] = base
	}
	target %= n
	for i := range target {
		ans[i]++
	}
	return ans
}
