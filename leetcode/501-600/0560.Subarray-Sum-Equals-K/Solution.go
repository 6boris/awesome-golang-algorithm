package Solution

func subarraySum(nums []int, k int) int {
	ans := 0
	for right := 0; right < len(nums); right++ {
		sum := 0
		for left := right; left >= 0; left-- {
			sum += nums[left]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}

func subarraySum_2(nums []int, k int) int {
	ans, sum, m := 0, 0, map[int]int{
		0: 1,
	}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := m[sum-k]; ok {
			ans += m[sum-k]
		}
		m[sum] += 1
	}

	return ans
}
