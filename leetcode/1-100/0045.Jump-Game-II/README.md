# [45. Jump Game II][title]

## Description

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.
**Example 1:**

```
Input: [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2.
    Jump 1 step from index 0 to 1, then 3 steps to the last index.
```


> Note:You can assume that you can always reach the last index.

**Tags:** Math, String

## 题意
>给定一个非负整数数组，你最初位于数组的第一个位置。
 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
## 题解

### 思路1
> DP 基本上会超时

```go
func jump(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < len(nums); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 0; i < n-1; i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < n {
				dp[i+j] = Min(dp[i+j], dp[i]+1)
			}
		}
	}
	return dp[n-1]
}
```

### 思路2
> 滑动窗口

```go
func jump(nums []int) int {
	//特殊情况判断
	if len(nums)==1{
		return 0
	}else if nums[0] >= len(nums){
		return 1
	}
	left, right, res := 0, nums[0], 0
	for right < len(nums)-1 {
		max := 0
		for i := left; i <= right; i++ {
			//寻找最大跨度
			if nums[i]-(right-i) >= max{
				max = nums[i]-(right - i)
			}
		}
		//窗口滑动右滑
		left = right
		right += max
		res++
	}
	//判断最后一步是不是踏在最后一个格子。循环条件结束只能确定能到达边界，left指针才是每次跳的格子。
	if left<len(nums)-1{
		res++
	}
	return res
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/jump-game-ii/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
