package Solution

func Solution(nums []int) []int {
	ret := make([]int, len(nums))
	for i, n := range nums {
		if n == 2 {
			ret[i] = -1
			continue
		}
		k, t := 0, n
		for ; t&1 == 1; t, k = t>>1, k+1 {

		}

		ret[i] = n ^ (1 << (k - 1))
	}
	return ret
}
