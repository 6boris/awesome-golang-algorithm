package Solution

func Solution(nums []int) int64 {
	l := len(nums)
	diff := make(map[int]int64)
	diff[0-nums[0]] = 1

	var ans, count int64
	count = 1
	for i := 1; i < l; i++ {
		a := i - nums[i]
		ans += count - diff[a]
		diff[a]++
		count++
	}
	return ans
}
