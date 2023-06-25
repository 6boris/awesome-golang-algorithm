package Solution

const mod1575 = 1000000007

func Solution(locations []int, start int, finish int, fuel int) int {
	// 需要dp[i][j] 表示 从第i个城市，还有fuel油量的时候能走的路径
	dp := make([][]int, len(locations))
	for i := 0; i < len(locations); i++ {
		dp[i] = make([]int, fuel+1)
		// 草，第一版这个dp没有做-1的初始化，直接超时，后来看了题解才想起来
		// 无路可走的时候，dp是等于0的。尴尬
		for j := 0; j <= fuel; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(_start, _fuel int) int {
		if _fuel < 0 {
			return 0
		}

		if dp[_start][_fuel] != -1 {
			return dp[_start][_fuel]
		}
		c := 0
		if _start == finish {
			c = 1
		}
		for i := 0; i < len(locations); i++ {
			if i == _start {
				continue
			}
			diff := locations[i] - locations[_start]
			if diff < 0 {
				diff = -diff
			}
			c = (c + dfs(i, _fuel-diff)) % mod1575
		}
		dp[_start][_fuel] = c
		return c
	}
	return dfs(start, fuel)
}
