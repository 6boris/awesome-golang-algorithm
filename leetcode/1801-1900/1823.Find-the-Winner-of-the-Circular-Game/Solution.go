package Solution

func Solution(n int, k int) int {
	winner := 0
	for i := 2; i <= n; i++ {
		winner = (winner + k) % i
	}
	return winner + 1
}
