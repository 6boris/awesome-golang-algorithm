package Solution

func countOfOne477(x int, v *[32]int) {
	index := 31
	for ; x > 0; index-- {
		if x&1 == 1 {
			v[index]++
		}
		x >>= 1
	}
}

func Solution(nums []int) int {
	ans := 0
	l := len(nums)
	ones := [32]int{}
	for i := 0; i < l; i++ {
		countOfOne477(nums[i], &ones)
	}
	for i := 0; i < 32; i++ {
		if ones[i] == 0 {
			continue
		}

		ans += ones[i] * (l - ones[i])
	}
	return ans
}
