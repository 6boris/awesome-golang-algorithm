package Solution

func Solution(nums []int) int {
	group := make(map[int][]int)
	for i := range nums {
		group[nums[i]] = append(group[nums[i]], i)
	}
	ret := -1
	for key := range group {
		if len(group[key]) < 3 {
			continue
		}
		for end := 2; end < len(group[key]); end++ {
			v := 2 * (group[key][end] - group[key][end-2])
			if ret == -1 || ret > v {
				ret = v
			}
		}
	}
	return ret
}
