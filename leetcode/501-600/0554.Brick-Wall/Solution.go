package Solution

func leastBricks(wall [][]int) int {
	m := make(map[int]int)
	max := 0
	for _, v := range wall {
		r := 0
		v = v[:len(v)-1]
		for _, vv := range v {
			m[r+vv]++
			if m[r+vv] > max {
				max = m[r+vv]
			}
			r += vv
		}
	}
	return len(wall) - max
}
