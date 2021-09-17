package Solution

func checkPerfectNumber(num int) bool {
	if num == 0 {
		return false
	}
	sum := 0

	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}
