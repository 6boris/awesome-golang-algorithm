# [300. Longest Increasing Subsequence][title]

## Description

Given an unsorted array of integers, find the length of longest increasing subsequence.
**Example 1:**

```
Input: [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
```

> Note:

- There may be more than one LIS combination, it is only necessary for you to return the length.
- Your algorithm should run in O(n2) complexity.

**Tags:** Math, String

## 题意
>给定一个无序的整数数组，找到其中最长上升子序列的长度。

## 题解

### 思路1
> （动态规划）O(n2)

- 用数组`dp[i]`记录以`nums[i]`结尾（即`nums[i]`为最后一个数字）的最长递增子序列的长度，则递推方程为`dp[i]=max(dp[j]+1)`，其中要求`1≤j<i`且`nums[j]<nums[i]`。

> 时间复杂度分析：

- 对每个i(1≤i≤n)，都需要从1遍历到i，则时间复杂度为O(n2)，空间复杂度的话需要一个额外的dp数组，空间复杂度为O(n2)。

> Code
```go
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp, ans := make([]int, len(nums)), 1

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			//fmt.Println(i, j, ans, dp, nums)
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
