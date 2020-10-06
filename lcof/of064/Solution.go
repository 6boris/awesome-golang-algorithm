package Solution

func sumNums(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumNums(n-1)
}
