package Solution

func Solution(nums []int, k int) int {
	const MOD = 1000000007
	cnt := make([]int, 30)

	// Count the number of 1 bits at each position
	for _, num := range nums {
		for i := 0; i < 30; i++ {
			if num&(1<<i) != 0 {
				cnt[i]++
			}
		}
	}

	res := int64(0)
	for i := 0; i < k; i++ {
		cur := int64(0)
		// Construct the current number from the most significant bit to the
		// least significant bit
		for j := 29; j >= 0; j-- {
			if cnt[j] > 0 {
				cur |= int64(1) << j
				cnt[j]--
			}
		}
		res = (res + cur*cur) % MOD
	}

	return int(res)
}
