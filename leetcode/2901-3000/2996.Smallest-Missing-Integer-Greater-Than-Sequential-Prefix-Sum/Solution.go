package Solution

func Solution(nums []int) int {
	sum := nums[0]
	exists := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		exists[nums[i]] = struct{}{}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			break
		}
		sum += nums[i]
	}

	for ; ; sum++ {
		_, ok := exists[sum]
		if !ok {
			break
		}
	}
	return sum
}
