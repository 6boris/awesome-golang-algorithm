package Solution

func Solution(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i = i + 2 {
		freq, val := nums[i], nums[i+1]
		for j := 0; j < freq; j++ {
			res = append(res, val)
		}
	}
	return res
}
