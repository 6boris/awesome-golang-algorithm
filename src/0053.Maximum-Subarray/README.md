# [1. Add Sum][title]

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
在遍历数组时用 Stack 把数组中的数存起来，如果当前遍历的数比栈顶元素来的大，说明栈顶元素的下一个比它大的数就是当前元素。
```go
func maxSubArray(nums []int) int {
	dp := nums[0]
	max := dp

	for i:=1; i<len(nums); i++ {
		if dp >0 {
			dp = nums[i] + dp
		} else {
			dp = nums[i]
		}
		if dp > max {
			max = dp
		}
	}

	return max
}
```
### 思路2


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/maximum-subarray/
[me]: https://github.com/kylesliu/awesome-golang-leetcode

