package Solution

func printNumbers(n int) []int {
	sum, ans := 1, []int{}
	for i := 0; i < n; i++ {
		sum *= 10
	}
	for i := 0; i < sum-1; i++ {
		ans = append(ans, i+1)
	}
	return ans
}

func printNumbers2(n int) []int {
	sum, ans := 1, []int{}
	for i := 0; i < n; i++ {
		sum *= 10
	}
	for i := 0; i < sum-1; i++ {
		ans = append(ans, i+1)
	}
	return ans
}
