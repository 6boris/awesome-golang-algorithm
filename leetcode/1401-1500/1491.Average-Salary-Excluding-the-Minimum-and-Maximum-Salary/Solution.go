package Solution

func Solution(salary []int) float64 {
	sum := 0
	m, n := -1, -1
	for _, nn := range salary {
		sum += nn
		if nn > m {
			m = nn
		}
		if n == -1 || nn < n {
			n = nn
		}
	}
	return float64(sum-m-n) / float64((len(salary) - 2))
}
