package Solution

func Solution(n int) int {
	digit := [10]int{}
	var ret int
	for n > 0 {
		digit[n%10]++
		n /= 10
	}
	for i := range 10 {
		ret += i * digit[i]
	}
	return ret
}
