package Solution

func Solution(low int, high int) int {
	numCount := high - low + 1
	x := numCount / 2
	if numCount&1 == 0 {
		return x
	}
	if low&1 == 1 {
		x++
	}
	return x
}
