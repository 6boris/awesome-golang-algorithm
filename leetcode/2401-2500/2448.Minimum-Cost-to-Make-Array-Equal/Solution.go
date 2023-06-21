package Solution

func cost2448(nums []int, cost []int, target int) int64 {
	c := int64(0)
	for idx, num := range nums {
		diff := num - target
		if diff < 0 {
			diff = ^diff + 1
		}
		c += int64(diff) * int64(cost[idx])
	}
	return c
}

// 这道题和之前的一个相似，忘了， 类似火车运送货物的
// 也是找到最大最大值，然后二分搜索
func Solution(nums []int, cost []int) int64 {
	max, min := -1, -1
	for _, i := range nums {
		if i > max {
			max = i
		}
		if min == -1 || i < min {
			min = i
		}
	}
	ans := int64(-1)
	for min < max {
		mid := (max-min)/2 + min
		c1 := cost2448(nums, cost, mid)
		c2 := cost2448(nums, cost, mid+1)
		ans = c1
		if c2 < c1 {
			ans = c2
		}
		if c2 < c1 {
			min = mid + 1
			continue
		}
		max = mid
	}
	return ans
}
