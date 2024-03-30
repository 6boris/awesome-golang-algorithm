package Solution

func Solution(nums []int, k int) int {
	ll := len(nums)
	count := make([]int, ll+1)
	ans := 0
	l, r := 0, 0
	cc := 0
	for ; r < ll; r++ {
		count[nums[r]]++
		if count[nums[r]] == 1 {
			k--
		}
		if k < 0 {
			count[nums[l]]--
			l++
			k++
			cc = 0
		}
		if k == 0 {
			for count[nums[l]] > 1 {
				count[nums[l]]--
				l++
				cc++
			}
			ans += (cc + 1)
		}
	}
	return ans
}
