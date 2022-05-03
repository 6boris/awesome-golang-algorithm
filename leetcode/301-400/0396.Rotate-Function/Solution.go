package Solution

func Solution(nums []int) int {
	/*
		timeout
		length := len(nums)
		if length <= 1 {
			return 0
		}
		ans := -2147483648
		for loop := 0; loop < len(nums); loop++ {
			base, tmpAns := 1, 0
			start := (length-loop)%length
			for next := (start+1)%length; next != start; next = (next+1)%length {
				tmpAns += (base*nums[next])
				base ++
			}
			if ans < tmpAns {
				ans = tmpAns
			}
		}
		return ans
	*/
	sum, ans := 0, 0
	for idx, n := range nums {
		sum += n
		ans += idx * n
	}

	maxAns := ans
	length := len(nums)
	// 0: 0*4 + 1*3 + 2*2 + 3*6 = 25
	// 1: 0*6 + 1*4 + 2*3 + 3*2 == > ans[0]-3*6 + 4+3+2 === > ans[i] = ans[n-i] - (n-1)(nums[n-i]) + sum-(n-i) = 16
	// 2: 0*2 + 1*6 + 2*4 + 3*3 == > ans[1]-3*2 + 6+4+3 = 23
	// ...
	for idx := 1; idx < length; idx++ {
		ans = ans - (length-1)*(nums[length-idx]) + sum - nums[length-idx]
		if ans > maxAns {
			maxAns = ans
		}
	}

	return maxAns
}
