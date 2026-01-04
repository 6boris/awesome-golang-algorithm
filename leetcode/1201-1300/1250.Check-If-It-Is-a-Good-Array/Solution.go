package Solution

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Solution(nums []int) bool {
	g := nums[0]
	for i := 1; i < len(nums); i++ {
		g = gcd(g, nums[i])
		if g == 1 {
			return true
		}
	}
	return g == 1
}
