package Solution

func Solution(nums []int, k int) int {
	start, end := 0, 0
	ans := 1
	length := len(nums)
	count := make(map[int]int)
	for ; end < length; end++ {
		count[nums[end]]++
		if count[nums[end]] <= k {
			if r := end - start + 1; r > ans {
				ans = r
			}
			continue
		}
		for count[nums[end]] > k {
			start++
			count[nums[start-1]]--
		}
	}
	return ans
}
