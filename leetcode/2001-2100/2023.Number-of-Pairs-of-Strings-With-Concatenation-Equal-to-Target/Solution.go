package Solution

func Solution(nums []string, target string) int {
	count := make(map[string]int)
	for _, s := range nums {
		count[s]++
	}
	ans := 0
	for i := 1; i < len(target); i++ {
		a, b := target[:i], target[i:]
		if a != b {
			ans += count[a] * count[b]
			continue
		}
		ans += count[a] * (count[a] - 1)
	}
	return ans
}
