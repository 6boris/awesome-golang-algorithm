# [561. Array Partition I][title]

## Description

Given an array of 2n integers, your task is to group these integers into n pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.
**Example 1:**

```
Input: [1,4,3,2]

Output: 4
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).

Note:
n is a positive integer, which is in the range of [1, 10000].
All the integers in the array will be in the range of [-10000, 10000].
```
**Tags:** Math, String

## 题意
> 给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1,b1),(a2,b2),…,(an,bn)(a1,b1),(a2,b2),…,(an,bn) ，
>使得从 1 到 n 的 min(ai,bi)min(ai,bi) 总和最大

## 题解

### 思路1
> 很容易想到，让两个比较小的数在一起比一个小的一个大的数放在一起产生的结果要大。
  故可以先将数组排序，然后从头开始相邻两个整数放在一起，这样子可以产生最大的结果。

- 时间复杂度
对数组进行排序，故时间复杂度为 O(nlogn)。

```go
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
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

[title]: https://leetcode.com/problems/array-partition-i/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
