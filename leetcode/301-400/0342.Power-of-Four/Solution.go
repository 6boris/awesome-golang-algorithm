package Solution

func Solution(n int) bool {
	if n < 1 {
		return false
	}
	// 4^2 == 2^4  4^n=2^(2n)
	// 所以需要先判断是否是2的幂次，然后判断这个幂次是不是偶数
	//  0101(8组, int 32) 55555555
	x := 0x55555555
	powerOf2 := n&(n-1) == 0
	if !powerOf2 {
		return false
	}
	return n&x != 0
}
