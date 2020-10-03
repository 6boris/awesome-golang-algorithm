package Solution

func isPalindrome(x int) bool {
	//	当x为负数，10的倍数时不可能是回文数
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	halfReverseX := 0
	for ; x > halfReverseX; x /= 10 {
		halfReverseX = halfReverseX*10 + x%10

	}
	//	需要排除数字个数为单数的时候
	return x == halfReverseX || halfReverseX/10 == x
}

func isPalindrome1(x int) bool {
	//	当x为负数，10的倍数时不可能是回文数
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	return x == reverse(x)
}

func reverse(x int) int {
	res := 0
	for ; x != 0; x /= 10 {
		res = res*10 + x%10
	}
	return res
}
