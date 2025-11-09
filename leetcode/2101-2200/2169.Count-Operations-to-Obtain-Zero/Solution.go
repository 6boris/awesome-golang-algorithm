package Solution

func Solution(num1 int, num2 int) int {
	var ret, x int
	for num1 > 0 && num2 > 0 {
		if num1 > num2 {
			x = num1 / num2
			ret += x
			num1 -= num2 * x
			continue
		}
		x = num2 / num1
		ret += x
		num2 -= num1 * x
	}
	return ret
}
