package Solution

func Solution(n int, x int) int64 {
	ans := int64(x)
	if n == 1 {
		return ans
	}
	nn := int64(n) - 1
	zero := 0
	c := int64(0)
	for {
		zero++
		c = 1 << (zero - 1)
		if nn <= c {
			break
		}
		nn -= c
	}
	c |= (nn - 1)
	cur := int64(1)
	for ; c > 0; c >>= 1 {
		for ans&cur != 0 {
			cur <<= 1
		}
		now := c & 1
		if now == 1 {
			ans |= cur
			continue
		}
		cur <<= 1
	}
	return ans
}
