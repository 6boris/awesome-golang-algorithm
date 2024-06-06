package Solution

func Solution(s string) bool {
	// 尝试将*替换成 “”，(, )，来做判断, 就是一个dfs+cache，但是问题是，*可能会非常多
	// 例如******* 一共100个，每个搜有三种选择，就离谱了啊, 先试试
	l := len(s)
	cache := make([][]int8, l+1)
	for i := 0; i <= l; i++ {
		cache[i] = make([]int8, l+1)
		for j := 0; j <= l; j++ {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) bool
	dfs = func(index, left int) bool {
		if index == l {
			return left == 0
		}

		if cache[index][left] != -1 {
			return cache[index][left] == 1
		}

		r := false
		if s[index] == '*' {
			// 如果 将* 设置为(
			r = r || dfs(index+1, left+1)
			// 注意只有说有(的时候，设置为)才是正确的
			if left > 0 {
				r = r || dfs(index+1, left-1)
			}
			r = r || dfs(index+1, left)
		} else {
			if s[index] == '(' {
				r = dfs(index+1, left+1)
			} else {
				if left > 0 {
					r = dfs(index+1, left-1)
				}
			}
		}
		if r {
			cache[index][left] = 1
		} else {
			cache[index][left] = 0
		}
		return r
	}
	return dfs(0, 0)
}
