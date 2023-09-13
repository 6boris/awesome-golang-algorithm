package Solution

func Solution(nums []int) []int {
	l := len(nums)
	neg := make([]int, l/2)
	pos := make([]int, l/2)
	ans := make([]int, l)
	ai, bi := 0, 0
	for _, n := range nums {
		if n < 0 {
			neg[ai] = n
			ai++
			continue
		}
		pos[bi] = n
		bi++
	}
	for i := 0; i < l; i += 2 {
		ans[i] = pos[i/2]
		ans[i+1] = neg[i/2]
	}
	return ans
}
