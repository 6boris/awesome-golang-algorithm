package Solution

func Solution(n int) int {
	var dfs func(int) int
	dfs = func(x int) int {
		if x == 1 {
			return 0
		}
		if x == 2 {
			return 1
		}
		odd := x&1 == 1
		if !odd {
			return dfs(x/2) + 1
		}
		x1 := dfs(x - 1)
		x2 := dfs(x + 1)
		if x1 < x2 {
			return x1 + 1
		}
		return x2 + 1
	}
	return dfs(n)
}
