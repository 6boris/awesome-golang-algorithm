package Solution

func Solution(nums []int, target int) int {
	cache := map[int]struct{}{
		0: {},
	}
	var ret int
	sum, need := 0, 0
	for i := range nums {
		sum += nums[i]
		need = sum - target
		if _, ok := cache[need]; ok {
			ret++
			cache = map[int]struct{}{
				0: {},
			}
			sum = 0
			continue
		}
		cache[sum] = struct{}{}
	}
	return ret
}
