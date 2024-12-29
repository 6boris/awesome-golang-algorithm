package Solution

func Solution(nums []int, k int) []int {
	l := len(nums)
	sl := l - k + 1
	sum := make([]int64, sl)
	cur := int64(0)
	i := 0
	for ; i < k; i++ {
		cur += int64(nums[i])
	}
	sum[0] = cur
	for ; i < l; i++ {
		cur -= int64(nums[i-k])
		cur += int64(nums[i])
		sum[i-k+1] = cur
	}
	maxIndex := make([]int, sl)
	maxIndex[sl-1] = sl - 1
	for i := sl - 2; i >= 0; i-- {
		maxIndex[i] = maxIndex[i+1]
		if sum[i] >= sum[maxIndex[i+1]] {
			maxIndex[i] = i
		}
	}
	m := int64(0)
	ans := make([]int, 3)
	for i := 0; i < l; i++ {
		for j := i + k; j < sl-k; j++ {
			s := sum[i] + sum[j] + sum[maxIndex[j+k]]
			if s > m {
				ans[0], ans[1], ans[2] = i, j, maxIndex[j+k]
				m = s
			}
		}
	}

	return ans
}
