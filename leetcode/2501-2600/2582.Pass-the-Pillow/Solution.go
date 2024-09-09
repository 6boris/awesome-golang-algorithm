package Solution

func Solution(n int, time int) int {
	need := n - 1

	loop := time / need
	left := time % need
	if left != 0 {
		loop++
	}

	if left == 0 {
		left = need
	}
	if loop&1 != 0 {
		return left + 1
	}
	return need - left + 1
}
