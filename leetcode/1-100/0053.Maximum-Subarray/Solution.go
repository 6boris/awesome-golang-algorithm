package Solution

func maxSubArray(nums []int) int {
	dp, ans := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
		ans = max(dp[i], ans)
		// fmt.Println(i, ans, dp, nums)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray2(nums []int) int {
	return calc(0, len(nums)-1, nums)
}

func calc(l, r int, nums []int) int {
	if l == r {
		return nums[l]
	}
	mid := (l + r) >> 1
	lmax, lsum, rmax, rsum := nums[mid], 0, nums[mid+1], 0

	for i := mid; i >= 1; i-- {
		lsum += nums[i]
		lmax = max(lmax, lsum)
	}

	for i := mid + 1; i <= r; i++ {
		rsum += nums[i]
		rmax = max(rmax, rsum)
	}

	return max(max(calc(l, mid, nums), calc(mid+1, r, nums)), lmax+rmax)
}
