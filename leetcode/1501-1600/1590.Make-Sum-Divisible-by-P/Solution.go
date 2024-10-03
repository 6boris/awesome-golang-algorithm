package Solution

func Solution(nums []int, p int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum = (sum + num) % p
	}

	target := sum % p
	if target == 0 {
		return 0
	}

	cache := make(map[int]int)
	cache[0] = -1
	tmpSum := 0
	ans := n

	for i := 0; i < n; i++ {
		tmpSum = (tmpSum + nums[i]) % p
		needed := (tmpSum - target + p) % p

		if idx, found := cache[needed]; found {
			ans = min(ans, i-idx)
		}

		cache[tmpSum] = i
	}

	if ans == n {
		return -1
	}
	return ans
}
