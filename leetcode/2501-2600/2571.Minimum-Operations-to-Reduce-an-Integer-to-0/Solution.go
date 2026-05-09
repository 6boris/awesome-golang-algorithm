package Solution

func Solution(n int) int {
	// 2个1的情况，直接点掉2个1与先+1 在减去高位1的步骤一样
	// 但是多个连续1，肯定是+1 比点掉多个1更快了
	// 在讨论不连续情况那就01，直接去掉就可以了
	ret := 0
	for ; n > 0; n >>= 1 {
		if n&1 == 0 {
			continue
		}
		// 连续多个的情况 >= 2
		if n&3 == 3 {
			ret++
			n++
			continue
		}
		ret++
		n--
	}
	return ret
}
