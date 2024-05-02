package Solution

func Solution(nums []int) int {
	e := make(map[int]struct{})
	ans := -1
	for _, n := range nums {
		if _, ok := e[-n]; ok {
			if n < 0 {
				n = -n
			}
			ans = max(ans, n)
			continue
		}
		e[n] = struct{}{}
	}
	return ans
}
