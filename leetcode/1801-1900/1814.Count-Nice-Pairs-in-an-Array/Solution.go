package Solution

const mod1814 = 1000000007

func convert1814(n int) int {
	base := 0
	for n > 0 {
		x := n % 10
		n /= 10
		base = base*10 + x
	}
	return base
}

// num[i] + rev(nums[j]) = nums[j] + rev(nums[i])
// nums[i]- rev(nums[i]) = nums[j] - rev(nums[j])
func Solution(nums []int) int {
	count := make(map[int]int)
	for _, n := range nums {
		x := n - convert1814(n)
		count[x]++
	}
	ans := 0
	for _, c := range count {
		ans = (ans + (c-1)*c/2) % mod1814
	}
	return ans
}
