package Solution

const (
	push = "Push"
	pop  = "Pop"
)

func Solution(target []int, n int) []string {
	ans := make([]string, 0)
	start := 0
	for input := 1; input <= n && start < len(target); input++ {
		if target[start] > input {
			ans = append(ans, push, pop)
			continue
		}
		ans = append(ans, push)
		start++
	}
	return ans
}
