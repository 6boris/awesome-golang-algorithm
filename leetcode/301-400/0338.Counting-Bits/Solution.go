package Solution

func Solution(n int) []int {
	res := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + 1
		}
	}
	return res
}
