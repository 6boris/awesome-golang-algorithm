package Solution

func Solution(nums []int, k int) int {
	cnt := map[int]int{}
	for i := range nums {
		if k == len(nums) {
			cnt[nums[i]] = 1
			continue
		}

		cnt[nums[i]]++
		if k == 1 || i == 0 || i == len(nums)-1 {
			continue
		}
		cnt[nums[i]] += 1
	}
	ret := -1
	for i := range nums {
		if cnt[nums[i]] == 1 {
			ret = max(ret, nums[i])
		}
	}
	return ret
}
