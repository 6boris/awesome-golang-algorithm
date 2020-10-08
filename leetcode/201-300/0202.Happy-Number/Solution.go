package Solution

//	快慢指针
func isHappy(n int) bool {
	slow, fast := n, squareSum(n)
	for fast != 1 && slow != fast {
		slow = squareSum(slow)
		fast = squareSum(squareSum(fast))
	}
	return fast == 1
}

//	利用 map 判断
func isHappy2(n int) bool {
	seen := make(map[int]int)
	for {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = -1
		n = squareSum(n)
		if n == 1 {
			return true
		}
	}
}

func squareSum(m int) int {
	sum := 0
	for m != 0 {
		sum += (m % 10) * (m % 10)
		m /= 10
	}
	return sum
}
