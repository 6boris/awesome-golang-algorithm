package Solution

func Solution(nums []int, key int, k int) []int {
	l := len(nums)
	used := make([]bool, l)
	c := 0
	for i := 0; i < l; i++ {
		if c > 0 {
			used[i] = true
			if nums[i] == key {
				c = k
			} else {
				c--
			}
			continue
		}
		if nums[i] == key {
			c = k
			used[i] = true
		}
	}
	c = 0
	for i := l - 1; i >= 0; i-- {
		if c > 0 {
			used[i] = true
			if nums[i] == key {
				c = k
			} else {
				c--
			}
			continue
		}
		if nums[i] == key {
			c = k
			used[i] = true
		}
	}
	var ans []int
	for i := range used {
		if used[i] {
			ans = append(ans, i)
		}
	}
	return ans
}
