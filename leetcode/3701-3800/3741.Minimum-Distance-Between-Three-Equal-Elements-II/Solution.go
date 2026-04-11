package Solution

func Solution(nums []int) int {
	found := false
	ret := 0x7fffffff
	group := make(map[int][]int)
	for i := range nums {
		group[nums[i]] = append(group[nums[i]], i)
	}
	for _, v := range group {
		if len(v) < 3 {
			continue
		}
		found = true
		for i := 2; i < len(v); i++ {
			ret = min(ret, 2*(v[i]-v[i-2]))
		}
	}
	if !found {
		return -1
	}
	return ret
}
