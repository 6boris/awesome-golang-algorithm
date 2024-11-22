package Solution

func Solution(n int) bool {
	checker := [10]int{}
	var checkNum func(n int)
	checkNum = func(n int) {
		for n > 0 {
			x := n % 10
			checker[x]++
			n /= 10
		}
	}
	checkNum(n)
	checkNum(n + n)
	checkNum(n + n + n)
	if checker[0] > 0 {
		return false
	}
	for i := 1; i < 10; i++ {
		if checker[i] != 1 {
			return false
		}
	}
	return true
}
