package Solution

func Solution(s string, k int) bool {
	sl := len(s)
	expect := pow2(k)
	if sl < k || int32(sl-k+1) < expect {
		return false
	}

	ans := make(map[int32]struct{})
	base := int32(0)
	m := int32(1)
	for idx := sl - 1; idx > sl-k-1; idx-- {
		base += int32(s[idx]-'0') * m
		m *= 2
	}
	m /= 2
	ans[base] = struct{}{}

	for idx := sl - 2; idx >= k-1; idx-- {
		base >>= 1
		base += m * int32(s[idx+1-k]-'0')
		ans[base] = struct{}{}
	}
	return int32(len(ans)) == expect
}

func pow2(k int) int32 {
	base := int32(1) // 2^0
	for i := 0; i < k; i++ {
		base *= 2
	}
	return base
}
