package Solution

func Solution(nums []int) int {
	one := 0
	for _, n := range nums {
		if n == 1 {
			one++
		}
	}
	l := len(nums)
	// 全是1, 0
	if one == l || one == 0 {
		return 0
	}
	// 0, 0, 1, 1, 0, 1
	// (0, 2), (1,3),(2,4),(3,5),(4,0),(5,1)
	// 0, 1, 2, 3, 4, 5
	start, end := l-one+1, 0
	ones := 0

	for i := start; i != end+1; i = (i + 1) % l {
		if nums[i] == 1 {
			ones++
		}
	}
	cur := ones
	if nums[start] == 1 {
		ones--
	}
	start, end = (start+1)%l, end+1
	for end < l {
		if nums[end] == 1 {
			ones++
		}
		cur = max(cur, ones)
		if nums[start] == 1 {
			ones--
		}
		start, end = (start+1)%l, end+1
	}

	return one - cur
}
