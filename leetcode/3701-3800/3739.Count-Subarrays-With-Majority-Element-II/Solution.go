package Solution

func Solution(nums []int, target int) int64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	offset := int64(n)
	count := make([]int64, 2*n+1)
	count[0+offset] = 1
	var ans int64 = 0
	curSum := int64(0)
	lessThanCount := int64(0)
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
