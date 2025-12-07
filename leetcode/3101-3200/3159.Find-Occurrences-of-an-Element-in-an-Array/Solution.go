package Solution

func Solution(nums []int, queries []int, x int) []int {
	indies := []int{}
	for idx := 0; idx < len(nums); idx++ {
		if nums[idx] == x {
			indies = append(indies, idx)
		}
	}
	ret := make([]int, len(queries))
	for idx, q := range queries {
		ret[idx] = -1
		if q <= len(indies) {
			ret[idx] = indies[q-1]
		}
	}
	return ret
}
