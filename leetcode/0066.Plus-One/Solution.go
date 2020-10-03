package Solution

func plusOne(digits []int) []int {
	// add one more slot to head for convenient
	digits = append([]int{0}, digits...)
	for i := len(digits) - 1; i >= 0; i-- {
		// find the first digit not 9 and plus one
		if digits[i] != 9 {
			digits[i] += 1
			break
		} else {
			// otherwise set to 0
			digits[i] = 0
		}
	}
	// if the first digit is 0 return the rest of slice
	if digits[0] == 0 {
		return digits[1:]
	}
	return digits
}

func Solution(digits []int) []int {
	n := len(digits)
	if n < 1 {
		return []int{1}
	}
	digits[n-1] += 1
	carry := 0
	for i := n - 1; i >= 0; i-- {
		n := digits[i] + carry
		digits[i], carry = n%10, n/10
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
