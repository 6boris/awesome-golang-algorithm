package Solution

func Solution(target []int) int {
	n := len(target)
	ans := target[0]
	for i := 1; i < n; i++ {
		ans += max(target[i]-target[i-1], 0)
	}
	return ans
}
