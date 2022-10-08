package Solution

func Solution(nums []int) bool {
	sumMap := make(map[int]struct{})
	for loop := 0; loop < len(nums)-1; loop++ {
		s := nums[loop] + nums[loop+1]
		if _, ok := sumMap[s]; ok {
			return true
		}
		sumMap[s] = struct{}{}
	}
	return false
}
