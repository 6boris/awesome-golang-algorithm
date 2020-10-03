# [162. Find Peak Element][title]

## Description

峰值定义为比左右相邻元素大的元素。
  给定一个数组 nums，保证 nums[i] ≠ nums[i+1]，请找出该数组的峰值，并返回峰值的下标。
  数组中可能包含多个峰值，只需返回任意一个即可。
  假定 nums[-1] = nums[n] = -∞。
  
**Example 1:**

```
输入：nums = [1,2,3,1]
输出：2
解释：3是一个峰值，3的下标是2。
```

**Example 2:**

```
输入：nums = [1,2,1,3,5,6,4]
输出：1 或 5 
解释：数组中有两个峰值：1或者5，返回任意一个即可。
```

**Tags:** Math, String

## 题意
> (二分) O(logn)O(logn)
  仔细分析我们会发现：

- 如果 `nums[i-1] < nums[i]`，则如果 `nums[i-1], nums[i], ... nums[n-1]` 是单调的，则 `nums[n-1]`就是峰值；如果`nums[i-1]`, `nums[i], ... nums[n-1]`不是单调的，则从 ii 开始，第一个满足 `nums[i] > nums[i+1]`的 ii 就是峰值；所以 `[i,n−1][i,n−1]` 中一定包含一个峰值；
- 如果 `nums[i-1] > nums[i]`，同理可得 `[0,i−1][0,i−1]` 中一定包含一个峰值；
 
所以我们可以每次二分中点，通过判断 `nums[i-1]` 和 `nums[i]` 的大小关系，可以判断左右两边哪边一定有峰值，从而可以将检索区间缩小一半。
  时间复杂度分析：二分检索，每次删掉一半元素，所以时间复杂度是 `O(logn)`。

## 题解

### 思路1
> 。。。。

```go

```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/find-peak-element/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
