package Solution

func Solution(nums []int) int {
	l := len(nums)
	sum := make([][2]int, l)
	if (l-1)&1 == 0 {
		sum[l-1][0] += nums[l-1]
	} else {
		sum[l-1][1] += nums[l-1]
	}
	for i := l - 2; i >= 0; i-- {
		sum[i] = sum[i+1]
		if i&1 == 0 {
			sum[i][0] += nums[i]
			continue
		}
		sum[i][1] += nums[i]
	}
	odd, even := 0, 0
	ans := 0
	for i := range l - 1 {
		if odd+sum[i+1][0] == even+sum[i+1][1] {
			ans++
		}
		if i&1 == 1 {
			odd += nums[i]
			continue
		}
		even += nums[i]
	}
	if odd == even {
		ans++
	}
	return ans
}
