package Solution

func Solution(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	offset := n
	count := make([]int, 2*n+1)
	count[offset] = 1
	var ans int = 0
	curSum := 0
	lessThanCount := 0
	for _, num := range nums {
		if num == target {
			lessThanCount += count[curSum+offset]
			curSum++
		} else {
			curSum--
			lessThanCount -= count[curSum+offset]
		}
		ans += lessThanCount
		count[curSum+offset]++
	}
	return ans
}
