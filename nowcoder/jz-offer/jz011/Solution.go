package Solution

func NumberOf1(val int) int {
	ans := 0
	for val != 0 {
		ans += 1
		val = val & (val - 1)
	}

	return ans
}
