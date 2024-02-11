package Solution

func Solution(nums []int, maximumBit int) []int {
	length := len(nums)
	ans := make([]int, length)
	xor := 0
	for _, n := range nums {
		xor ^= n
	}

	limit := 1 << maximumBit
	for i := 0; i < length; i++ {
		cur := 1 << 31
		curAns := 0
		for j := 31; j >= 0; j, cur = j-1, cur>>1 {
			if cur >= limit || cur&xor != 0 {
				continue
			}
			curAns |= cur
		}
		ans[i] = curAns
		xor ^= nums[length-1-i]
	}
	return ans
}
