package Solution

func Solution(n int) int {
	digits := [10]int{}
	for ; n > 0; n /= 10 {
		digits[n%10]++
	}
	ret, cnt := 1, 2
	for i := 9; i >= 0 && cnt > 0; i-- {
		if digits[i] == 0 {
			continue
		}
		digits[i]--
		cnt--
		ret *= i
		if digits[i] > 0 {
			i++
		}
	}
	return ret
}
