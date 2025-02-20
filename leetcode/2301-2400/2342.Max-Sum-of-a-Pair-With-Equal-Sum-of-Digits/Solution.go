package Solution

func countOfOne2342(n int) int {
	one := 0
	for n > 0 {
		one += n % 10
		n /= 10
	}
	return one
}

func Solution(nums []int) int {
	ans := -1
	sum := make(map[int][2]int)
	for i := range nums {
		one := countOfOne2342(nums[i])
		v, ok := sum[one]
		if !ok {
			sum[one] = [2]int{nums[i], -1}
			continue
		}
		if nums[i] >= sum[one][0] {
			v[1] = v[0]
			v[0] = nums[i]
			sum[one] = v
			continue
		}
		if nums[i] > sum[one][1] {
			v[1] = nums[i]
			sum[one] = v
		}
	}
	for _, v := range sum {
		if v[0] != -1 && v[1] != -1 {
			ans = max(ans, v[0]+v[1])
		}
	}
	return ans
}
