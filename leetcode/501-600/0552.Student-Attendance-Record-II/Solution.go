package Solution

const mod552 = 1000000007

func Solution(n int) int {
	// 直接dfs很明显，10101那个例子就过不去了,需要增加cache
	cache := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			cache[i][j] = make([]int, 4)
			for k := 0; k < 4; k++ {
				cache[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(nn, a, l int) int {
		// a<2&&l<3
		if a >= 2 || l >= 3 {
			return 0
		}
		if nn == 0 {
			//放满了，
			return 1
		}
		if cache[nn][a][l] != -1 {
			return cache[nn][a][l]
		}
		ans := 0
		// 选择一个P,下一个随便放
		ans = (ans + dfs(nn-1, a, 0)) % mod552
		//  放置一个A
		ans = (ans + dfs(nn-1, a+1, 0)) % mod552
		// 放置一个L, L需要连续，所以放A和P的时候将L设置为0，这里是+1
		ans = (ans + dfs(nn-1, a, l+1)) % mod552
		cache[nn][a][l] = ans
		return ans
	}
	return dfs(n, 0, 0)
}
