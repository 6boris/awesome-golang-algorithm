package Solution

func gcd2447(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Solution(nums []int, k int) int {
	size := len(nums)
	var ret int
	for i := 0; i < size; i++ {
		cur := nums[i]
		if cur == k {
			ret++
		}
		for j := i + 1; j < size; j++ {
			cur = gcd2447(cur, nums[j])
			if cur < k {
				break
			}
			if cur == k {
				ret++
			}
		}
	}
	return ret
}
