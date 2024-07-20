package Solution

func Solution(rowSum []int, colSum []int) [][]int {
	N := len(rowSum)
	M := len(colSum)

	currRowSum := make([]int, N)
	currColSum := make([]int, M)

	origMatrix := make([][]int, N)
	for i := 0; i < N; i++ {
		origMatrix[i] = make([]int, M)
		for j := 0; j < M; j++ {
			origMatrix[i][j] = min(rowSum[i]-currRowSum[i], colSum[j]-currColSum[j])

			currRowSum[i] += origMatrix[i][j]
			currColSum[j] += origMatrix[i][j]
		}
	}
	return origMatrix
}
