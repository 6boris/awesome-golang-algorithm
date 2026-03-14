package Solution

func Solution(n int) int {

	begin, init := 1, 1
	half, ret := n/2, 0
	for i := 0; i < n; i++ {
		ret++
		if init&1 == 0 {
			init /= 2
		} else {
			init = half + (init-1)/2
		}
		if init == begin {
			break
		}
	}
	return ret
}
