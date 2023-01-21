package Solution

func Solution(nums []int) [][]int {
	ans := make([][]int, 0)
	l := len(nums)
	visited := make(map[[15]int]struct{})
	var dfs func(int, [15]int, int, map[[15]int]struct{})
	dfs = func(start int, path [15]int, idx int, visited map[[15]int]struct{}) {
		if start >= l {
			return
		}
		for next := start + 1; next < l; next++ {
			if nums[next] < nums[start] {
				continue
			}
			path[idx] = nums[next]
			if _, ok := visited[path]; ok {
				continue
			}
			visited[path] = struct{}{}
			dst := make([]int, idx+1)
			copy(dst, path[:idx+1])
			ans = append(ans, dst)
			dfs(next, path, idx+1, visited)
		}
	}
	for i := 0; i < l-1; i++ {
		path := [15]int{nums[i], -101, -101, -101, -101, -101, -101, -101, -101, -101, -101, -101, -101, -101, -101}
		dfs(i, path, 1, visited)
	}
	return ans
}
