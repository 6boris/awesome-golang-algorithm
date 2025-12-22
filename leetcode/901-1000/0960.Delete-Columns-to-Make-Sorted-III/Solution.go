package Solution

func Solution(A []string) int {
	if len(A) == 0 {
		return 0
	}

	numCols := len(A[0])
	dp := make([]int, numCols)

	for i := range dp {
		dp[i] = 1
	}

	for i := numCols - 2; i >= 0; i-- {
		for j := i + 1; j < numCols; j++ {
			isSorted := true
			for _, row := range A {
				if row[i] > row[j] {
					isSorted = false
					break
				}
			}

			if isSorted {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	maxKept := 0
	for _, val := range dp {
		if val > maxKept {
			maxKept = val
		}
	}

	return numCols - maxKept
}
