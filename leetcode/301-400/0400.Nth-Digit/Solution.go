package Solution

func Solution(n int) int {
	if n <= 9 {
		return n
	}
	fixed := 9
	base, pow := 1, 1
	cur := base * pow * fixed
	for cur < n {
		n -= cur
		base, pow = base+1, pow*10
		cur = base * pow * fixed
	}
	need := n / base
	target := pow + (need - 1)
	mod := n % base
	if mod == 0 {
		return target % 10
	}

	shift := base - mod
	target++
	for ; shift > 0; shift-- {
		target /= 10
	}
	return target % 10
}
