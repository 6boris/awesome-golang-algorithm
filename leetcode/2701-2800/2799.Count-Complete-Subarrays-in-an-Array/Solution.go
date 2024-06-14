package Solution

func Solution(nums []int) int {
	in := make(map[int]struct{})
	for _, n := range nums {
		in[n] = struct{}{}
	}
	distinct := len(in)
	ans := 0
	for i := 0; i < len(nums); i++ {
		ti := map[int]struct{}{nums[i]: {}}
		if len(ti) == distinct {
			ans++
		}
		for pre := i - 1; pre >= 0; pre-- {
			ti[nums[pre]] = struct{}{}
			if len(ti) == distinct {
				ans++
			}
		}
	}
	return ans
}
