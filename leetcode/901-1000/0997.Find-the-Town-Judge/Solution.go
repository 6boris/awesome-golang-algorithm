package Solution

func Solution(n int, trust [][]int) int {
	ab, ba := make([]int, n+1), make([]int, n+1)
	for _, t := range trust {
		a, b := t[0], t[1]
		ab[a]++
		ba[b]++
	}
	for p := 1; p <= n; p++ {
		if ab[p] == 0 && ba[p] == n-1 {
			return p
		}
	}
	return -1
}
