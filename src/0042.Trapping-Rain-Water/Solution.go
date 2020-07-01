package Solution

import "fmt"

// 三次线性扫描
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	n, ans := len(height), 0
	leftMax, rightMax := make([]int, n), make([]int, n)

	//	求第i个矩形左边最高矩形的高度
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	//	求第i个矩形右边最高矩形的高度
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i := 0; i < n; i++ {
		fmt.Println(i, leftMax[i], rightMax[i], height[i])
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
