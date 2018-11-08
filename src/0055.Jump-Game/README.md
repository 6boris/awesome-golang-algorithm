# [55. Jump Game][title]

## Description

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.
**Example 1:**

```
Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

**Example 2:**

```
Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.
```

**Tags:** Math, String

## 题意
>给定一个非负整数数组，你最初位于数组的第一个位置。
 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 判断你是否能够到达最后一个位置。

## 题解

### 思路1
>  DP的思路是创建一个DP的Array用来存储int
>  状态：在每个节点，比对之前的所有节点，若之前节点为1且之前节点拥有到达目前且超过目前index的能力
```go
func canJump(nums []int) bool {
	dp := make([]int, len(nums))
	dp[len(nums)-1] = 1

	for i := len(nums) - 2; i >= 0; i-- {
		for j := 0; j <= nums[i] && j < len(nums); j++ {
			if dp[i+j] == 1 {
				dp[i] = 1
				break
			}
		}
	}
	return dp[0] == 1
}
```

### 思路2
> Greedy 这个是借鉴LeetCode上某个大佬的[@sohammehta][sohammehta-solution]
```go
func canJump2(nums []int) bool {
	n, farset := len(nums), 0
	for i := 0; i < n; i++ {
		if farset < i {
			return false
		}
		farset = Max(i+nums[i], farset)
	}
	return true
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/jump-game/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
[sohammehta-solution]:https://leetcode.com/problems/jump-game/discuss/182034/Difference-Between-DP-and-Greedy