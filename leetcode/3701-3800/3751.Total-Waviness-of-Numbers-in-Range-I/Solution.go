package Solution

func Solution(num1 int, num2 int) int {
	var ret int
	waviness := func(n int) int {
		var a, b, c, cnt int
		a = n % 10
		n /= 10
		b = n % 10
		n /= 10

		for ; n > 0; n /= 10 {
			c = n % 10
			if (b > a && b > c) || (b < a && b < c) {
				cnt++
			}
			a, b = b, c
		}
		return cnt
	}
	for i := max(100, num1); i <= num2; i++ {
		ret += waviness(i)
	}
	return ret
}
