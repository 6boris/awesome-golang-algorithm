package Solution

func Solution(nums []int) int {
	sum := 0
	pre := 0
	ans := 0
	for _, n := range nums {
		if n > pre {
			sum += n
			pre = n
			continue
		}
		ans = max(ans, sum)
		sum = n
		pre = n
	}
	return max(ans, sum)
}
