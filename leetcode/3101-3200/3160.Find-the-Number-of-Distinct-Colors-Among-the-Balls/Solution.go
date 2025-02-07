package Solution

func Solution(limit int, queries [][]int) []int {
	ans := make([]int, len(queries))
	colors := make(map[int]int)
	colored := make(map[int]int)
	for i, q := range queries {
		ball, color := q[0], q[1]
		sourceColor, ok := colored[ball]
		if ok {
			colors[sourceColor]--
			if colors[sourceColor] == 0 {
				delete(colors, sourceColor)
			}
		}

		colors[color]++
		colored[ball] = color
		ans[i] = len(colors)
	}
	return ans
}
