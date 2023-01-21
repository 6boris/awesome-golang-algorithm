package Solution

const min918 = -30000

func Solution(nums []int) int {
	ans := min918
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	// 首先计算没有首尾相连的情况
	sub := make([]int, l)
	copy(sub, nums)
	for i := 0; i < l; i++ {
		if i != 0 {
			if sub[i-1]+sub[i] > sub[i] {
				sub[i] = sub[i-1] + sub[i]
			}
		}
		if sub[i] > ans {
			ans = sub[i]
		}
	}

	// lr记录到i时候，连续数组的最大和例如1,-1,2, lr[0] = 1, lr[1] = 1, lr[2] = 2
	// rl 同理
	// sum就是对数组加和，我们最后只需要对rl[i]+lr[i-1]与最大结果比较即可
	lr := make([]int, l)
	rl := make([]int, l)
	sum := make([]int, l)
	sum[0] = nums[0]
	lm := nums[0]
	for i := 1; i < l; i++ {
		sum[i] = sum[i-1] + nums[i]
		if sum[i] > lm {
			lm = sum[i]
		}
		lr[i] = lm
	}
	total := sum[l-1]
	if total > ans {
		ans = total
	}

	rm := nums[l-1]
	for i := l - 1; i > 0; i-- {
		sum[i] = total - sum[i-1]
		if rm < sum[i] {
			rm = sum[i]
		}
		rl[i] = rm
	}
	for i := l - 1; i > 0; i-- {
		if lr[i-1]+rl[i] > ans {
			ans = lr[i-1] + rl[i]
		}
	}
	return ans
}
