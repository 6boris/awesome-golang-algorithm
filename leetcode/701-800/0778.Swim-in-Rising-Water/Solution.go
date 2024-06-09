package Solution

func dfs778(grid [][]int, visited [][]bool, x, y, t int) bool {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid) || visited[x][y] || grid[x][y] > t {
		return false
	}
	if x == len(grid)-1 && y == len(grid)-1 {
		return true
	}
	visited[x][y] = true
	return dfs778(grid, visited, x-1, y, t) ||
		dfs778(grid, visited, x+1, y, t) ||
		dfs778(grid, visited, x, y-1, t) ||
		dfs778(grid, visited, x, y+1, t)

}
func do778(grid [][]int, x, y, t int) bool {
	gridCopy := make([][]int, len(grid))
	v := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		gridCopy[i] = make([]int, len(grid))
		v[i] = make([]bool, len(grid))
		copy(gridCopy[i], grid[i])
	}

	return dfs778(gridCopy, v, x, y, t)
}
func Solution(grid [][]int) int {

	// 一般这种dfs或者bfs都不太好确定的题，可以尝试binary search+dfs
	// 直接尝试搜索每个答案
	ans := -1
	l, r := 0, 1<<31-1
	for l <= r {
		mid := l + (r-l)/2
		if do778(grid, 0, 0, mid) {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}
