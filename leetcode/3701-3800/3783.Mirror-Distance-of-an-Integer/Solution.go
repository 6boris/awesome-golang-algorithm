package Solution

func Solution(n int) int {
	rev, src := 0, n
	for ; n > 0; n /= 10 {
		rev = rev*10 + n%10
	}
	diff := rev - src
	if diff < 0 {
		diff = -diff
	}
	return diff
}
