package Solution

var count int

func totalNQueens3(n int) int {
	if n < 1 {
		return 0
	}
	dfs(n, 0, 0, 0, 0)
	return count
}

func dfs(n int, row int, cols int, pie int, na int) {
	if row >= n {
		count++
		return
	}
	bits := (^(cols | pie | na) & ((1 << uint(n)) - 1)) // 查看当前能够放置元素的位置
	for bits > 0 {
		p := bits & -bits // 获取最低位1
		dfs(n, row+1, cols|p, (pie|p)>>1, (na|p)<<1)
		bits = bits & (bits - 1) // 去掉最低位1
	}
}
