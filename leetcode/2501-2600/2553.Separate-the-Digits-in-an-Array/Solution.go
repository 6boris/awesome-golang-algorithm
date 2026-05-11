package Solution

func Solution(nums []int) []int {
	ret := make([]int, 0)
	var n, tmp int
	for i := range nums {
		n = nums[i]
		var bits []int
		for n > 0 {
			tmp = n % 10
			bits = append(bits, tmp)
			n /= 10
		}
		for j := len(bits) - 1; j >= 0; j-- {
			ret = append(ret, bits[j])
		}
	}
	return ret
}
