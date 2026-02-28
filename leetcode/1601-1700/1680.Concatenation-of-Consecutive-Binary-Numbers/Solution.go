package Solution

const mod1680 = 1000000007

func Solution(n int) int {
	bits := func(x int) int {
		mask := 1 << 31
		cnt := 0
		for mask&x == 0 {
			mask >>= 1
			cnt++
		}
		return 32 - cnt
	}
	ret := 0
	for i := 1; i <= n; i++ {
		shift := bits(i)
		ret <<= shift
		ret %= mod1680
		ret = (ret + i) % mod1680
	}
	return ret
}
