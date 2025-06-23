package Solution

func Solution(k int, n int) int64 {
	digit := make([]int, 100)
	left, count, ans := 1, 0, int64(0)
	for count < n {
		right := left * 10
		for op := 0; op < 2; op++ {
			// enumerate i'
			for i := left; i < right && count < n; i++ {
				combined := int64(i)
				x := i
				if op == 0 {
					x = i / 10
				}
				for x > 0 {
					combined = combined*10 + int64(x%10)
					x /= 10
				}
				if isPalindrome(combined, k, digit) {
					count++
					ans += combined
				}
			}
		}
		left = right
	}
	return ans
}

func isPalindrome(x int64, k int, digit []int) bool {
	length := -1
	for x > 0 {
		length++
		digit[length] = int(x % int64(k))
		x /= int64(k)
	}
	for i, j := 0, length; i < j; i, j = i+1, j-1 {
		if digit[i] != digit[j] {
			return false
		}
	}
	return true
}
