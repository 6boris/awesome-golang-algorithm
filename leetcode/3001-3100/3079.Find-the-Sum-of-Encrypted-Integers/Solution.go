package Solution

func to(n int) int {
	x, bits, mod := 0, 0, 0
	for n > 0 {
		mod = n % 10
		n /= 10
		x = max(x, mod)
		bits++
	}
	base := 0
	for ; bits > 0; bits-- {
		base = base*10 + x
	}
	return base
}

func Solution(nums []int) int {
	var ret int
	for _, n := range nums {
		ret += to(n)
	}
	return ret
}
