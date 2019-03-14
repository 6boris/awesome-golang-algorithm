package Solution

//	普通的动态规划
func rob(nums []int) int {
	nLenghth := len(nums)

	dp := make([][2]int, nLenghth+1)
	for i := 1; i <= nLenghth; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = nums[i-1] + dp[i-1][0]
	}
	//fmt.Println(dp)
	return Max(dp[nLenghth][0], dp[nLenghth][1])
}

//	优化空间
func rob2(nums []int) int {
	prevNo, prevYes := 0, 0

	for _, v := range nums {
		tmp := prevNo
		prevNo = Max(prevNo, prevYes)
		prevYes = v + tmp
	}
	return Max(prevNo, prevYes)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func rob3(nums []int) int {
	robEven, robOdd := 0, 0
	for i := 0; i < len(nums); i++ {
		if i % 2 ==0 {
			robEven = Max(robEven + nums[i] , robOdd)
		} else {
			robOdd = Max(robEven , nums[i] + robOdd)
		}
	}
	return Max(robEven,robOdd)
}
