package Solution

func Solution(grid [][]int) int {
	used := map[[2]int]struct{}{}
	rows, cols := len(grid), len(grid[0])
	freshOranges := 0
	queue := [][2]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
				used[[2]int{r, c}] = struct{}{}
			}
			if grid[r][c] == 1 {
				freshOranges++
			}
		}
	}
	var dirs = [][2]int{
		{-1, 0}, {1, 0}, {0, 1}, {0, -1},
	}
	steps := 0
	for len(queue) > 0 {
		nq := make([][2]int, 0)
		for _, cur := range queue {
			for _, dir := range dirs {
				nx, ny := cur[0]+dir[0], cur[1]+dir[1]
				if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == 1 {
					key := [2]int{nx, ny}
					if _, ok := used[key]; !ok {
						nq = append(nq, key)
						used[key] = struct{}{}
					}
				}
			}
		}
		freshOranges -= len(nq)
		if len(nq) != 0 {
			steps++
		}
		queue = nq
	}
	if freshOranges > 0 {
		return -1
	}
	return steps
}
