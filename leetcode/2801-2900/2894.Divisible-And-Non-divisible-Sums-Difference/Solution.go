package Solution

func Solution(n int, m int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%m != 0 {
			sum += i
			continue
		}
		sum -= i
	}
	return sum
}
