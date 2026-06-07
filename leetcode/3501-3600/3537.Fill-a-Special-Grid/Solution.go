package Solution

func Solution(n int) [][]int {
	size := 1 << n
	ret := make([][]int, size)
	for i := range size {
		ret[i] = make([]int, size)
	}

	var dfs func(int, int, int, int)
	dfs = func(x, y, width, start int) {
		if width == 1 {
			ret[x][y] = start
			return
		}
		mid := width / 2
		nums := mid * mid

		// right top
		dfs(x, y, mid, start)
		// right bottom
		dfs(x+mid, y, mid, start+nums)
		// left bottom
		dfs(x+mid, y-mid, mid, start+nums+nums)
		// left top
		dfs(x, y-mid, mid, start+nums+nums+nums)
	}
	dfs(0, size-1, size, 0)
	return ret
}
