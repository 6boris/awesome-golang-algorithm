package Solution

func Solution(points [][]int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for idx, q := range queries {
		distance := q[2] * q[2]
		for _, p := range points {
			a := q[0] - p[0]
			b := q[1] - p[1]
			if a*a+b*b <= distance {
				ans[idx]++
			}
		}
	}
	return ans
}
