package Solution

func Solution(word string) int {
	pos := [26][2]int{}
	for i := range 26 {
		row, col := i/6, i%6
		pos[i] = [2]int{row, col}
	}

	distance := func(a, b int) int {
		if a == -1 || b == -1 {
			return 0
		}
		x, y := pos[a][0], pos[a][1]
		x1, y1 := pos[b][0], pos[b][1]
		lx := x - x1
		ly := y - y1
		if lx < 0 {
			lx = -lx
		}
		if ly < 0 {
			ly = -ly
		}
		return lx + ly
	}
	memo := make(map[[3]int]int)
	var dfs func(int, int, int) int
	dfs = func(left, right, index int) int {
		if index == len(word) {
			return 0
		}
		key := [3]int{left, right, index}
		if v, ok := memo[key]; ok {
			return v
		}

		cIndex := int(word[index] - 'A')
		leftCost := distance(left, cIndex) + dfs(cIndex, right, index+1)
		rightCost := distance(right, cIndex) + dfs(left, cIndex, index+1)
		memo[key] = min(leftCost, rightCost)
		return memo[key]
	}
	return dfs(-1, -1, 0)
}
