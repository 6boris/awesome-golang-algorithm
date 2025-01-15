package Solution

func countOfOne(x int) int {
	c := 0
	for x > 0 {
		c++
		x = x & (x - 1)
	}
	return c
}

func Solution(num1 int, num2 int) int {
	n1, n2 := countOfOne(num1), countOfOne(num2)
	if n1 == n2 {
		return num1
	}
	if n1 > n2 {
		diff := n1 - n2
		mask := 1
		for i := 0; i < 32 && diff > 0; i++ {
			if mask&num1 == mask {
				diff--
			}
			mask <<= 1
		}
		return num1 & ^(mask - 1)
	}
	ans := num1
	mask := 1
	diff := n2 - n1

	for i := 0; i < 32 && diff > 0; i++ {
		if mask&num1 == 0 {
			ans |= mask
			diff--
		}
		mask <<= 1
	}
	return ans
}
