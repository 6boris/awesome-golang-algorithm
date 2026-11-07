package Solution

func Solution(nums []int) int {
	f := make(map[int]int)
	for i := range nums {
		f[nums[i]]++
	}

	ff := make(map[int]int)
	for _, freq := range f {
		ff[freq]++
	}
	for i := range nums {
		if ff[f[nums[i]]] == 1 {
			return nums[i]
		}
	}

	return -1
}
