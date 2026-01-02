package Solution

func Solution(nums []int) int {
	in := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := in[n]; ok {
			return n
		}
		in[n] = struct{}{}
	}
	return -1
}
