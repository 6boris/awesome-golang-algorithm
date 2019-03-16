# [376. Wiggle Subsequence][title]

## Description

A sequence of numbers is called a wiggle sequence if the differences between successive numbers strictly alternate between positive and negative. The first difference (if one exists) may be either positive or negative. A sequence with fewer than two elements is trivially a wiggle sequence.

For example, [1,7,4,9,2,5] is a wiggle sequence because the differences (6,-3,5,-7,3) are alternately positive and negative. In contrast, [1,4,7,2,5] and [1,7,4,5,5] are not wiggle sequences, the first because its first two differences are positive and the second because its last difference is zero.

Given a sequence of integers, return the length of the longest subsequence that is a wiggle sequence. A subsequence is obtained by deleting some number of elements (eventually, also zero) from the original sequence, leaving the remaining elements in their original order.

**Example 1:**

```
Input: [1,7,4,9,2,5]
Output: 6
Explanation: The entire sequence is a wiggle sequence.
```

**Example 2:**

```
Input: [1,17,5,10,13,15,10,5,16,8]
Output: 7
Explanation: There are several subsequences that achieve this length. One is [1,17,10,13,10,16,8].
```

**Example 3:**

```
Input: [1,2,3,4,5,6,7,8,9]
Output: 2
```

**Tags:** Math, String

## 题意
>求摇摆序列最长序列，可以转化为数组【上升】【下降】状态的转换

<a href="https://github.com/kylesliu/awesome-golang-leetcode/tree/master/assets/images/376_Wiggle_Subsequence.jpg">
    <img src="https://github.com/kylesliu/awesome-golang-leetcode/tree/master/assets/images/376_Wiggle_Subsequence.jpg" />
</a>

## 题解

### 思路1
> 按照小学算数那么来做，用 `carry` 表示进位，从后往前算，依次往前，每算出一位就插入到最前面即可，直到把两个二进制串都遍历完即可。

```go
func wiggleMaxLength(nums []int) int {
	//	序列个数小于二直接是摇摆序列
	if len(nums) < 2 {
		return len(nums)
	}
	//	3种状态
	const BEGIN, UP, DOWN = 0, 1, 2
	//	摇摆序列最小长度至少为1
	STATE, maxLength := BEGIN, 1

	for i := 1; i < len(nums); i++ {
		switch STATE {
		case BEGIN:
			if nums[i-1] < nums[i] {
				STATE = UP
				maxLength++
			} else if nums[i-1] > nums[i] {
				STATE = DOWN
				maxLength++
			}
		case UP:
			if nums[i-1] > nums[i] {
				STATE = DOWN
				maxLength++
			}
		case DOWN:
			if nums[i-1] < nums[i] {
				STATE = UP
				maxLength++
			}
		}
	}
	return maxLength
}

```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/wiggle-subsequence/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
