package Solution

func Solution(encoded []int, first int) []int {
	ans := make([]int, len(encoded)+1)
	ans[0] = first
	//  1, 2, 3
	//1 x
	for idx, code := range encoded {
		ans[idx+1] = ans[idx] ^ code
	}
	return ans
}
