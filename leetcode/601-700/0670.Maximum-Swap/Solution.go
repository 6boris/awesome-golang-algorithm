package Solution

func Solution(num int) int {
	digits := make([]uint8, 0)
	for num > 0 {
		mod := num % 10
		digits = append(digits, uint8(mod))
		num /= 10
	}
	index := len(digits) - 1
	for ; index > 0; index-- {
		targetIndex := index
		for pre := index - 1; pre >= 0; pre-- {
			if digits[pre] >= digits[targetIndex] {
				targetIndex = pre
			}
		}
		if targetIndex != index && digits[targetIndex] != digits[index] {
			digits[targetIndex], digits[index] = digits[index], digits[targetIndex]
			break
		}
	}
	ans := 0
	for i := len(digits) - 1; i >= 0; i-- {
		ans = ans*10 + int(digits[i])
	}
	return ans
}
