package Solution

func Solution(queries []int, m int) []int {
	res, p := make([]int, len(queries)), make([]int, m)
	for i := 0; i < m; i++ {
		p[i] = i + 1
	}
	for idx, q := range queries {
		pos := -1
		for i, v := range p {
			if v == q {
				pos = i
				break
			}
		}
		res[idx] = pos
		for i := pos; i > 0; i-- {
			p[i] = p[i-1]
		}
		p[0] = q
	}
	return res
}
