package Solution

//	快慢指针
func isHappy_1(n int) bool {
	slow, fast := n, squareSum(n)
	for fast != 1 && slow != fast {
		slow = squareSum(slow)
		fast = squareSum(squareSum(fast))
	}
	return fast == 1
}

//	利用 map 判断
func isHappy_2(n int) bool {
	m := map[int]bool{}
	for n != 1 && !m[n] {
		m[n] = true
		n = squareSum(n)
	}
	return n == 1
}

func squareSum(m int) int {
	sum := 0
	for m != 0 {
		sum += (m % 10) * (m % 10)
		m /= 10
	}
	return sum
}
