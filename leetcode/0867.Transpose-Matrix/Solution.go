package Solution

func Solution(A [][]int) [][]int {
	if len(A) == len(A[0]) {
		for i := 0; i < len(A); i++ {
			for j := 0; j < len(A); j++ {
				if i == j {
					break
				}
				A[i][j], A[j][i] = A[j][i], A[i][j]
			}
		}
		return A
	}
	res := make([][]int, len(A[0]))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(A))
		for j := 0; j < len(A); j++ {
			res[i][j] = A[j][i]
		}
	}
	return res
}
