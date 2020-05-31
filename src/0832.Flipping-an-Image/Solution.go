package Solution

func Solution(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		n, r := len(A[i]), A[i]
		for j := 0; j < n/2; j++ {
			r[j], r[n-j-1] = 1-r[n-j-1], 1-r[j]
		}
		if n%2 == 1 {
			r[n/2] = 1 - r[n/2]
		}
	}
	return A
}
