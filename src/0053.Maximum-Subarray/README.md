# [53. Maximum Subarray][title]

## Description

Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

**Example:**

```
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

**Follow up:**

If you have figured out the O(*n*) solution, try coding another solution using the divide and conquer approach, which is more subtle.

**Tags:** Array, Divide and Conquer, Dynamic Programming


## 题解
### 思路1
> 普通DP 找到公式就好了
```$xslt
如果nums[0]>=0，则sum = sum + nums[1]
如果nums[0]<0，则sum = nums[1]，
```

```go
func maxSubArray(nums []int) int {
	dp, ans := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = Max(nums[i], nums[i]+dp[i-1])
		ans = Max(ans, dp[i])
	}

	return ans
}
```
### 思路2


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/maximum-subarray/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode

