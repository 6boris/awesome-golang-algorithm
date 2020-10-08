# [628. Maximum Product of Three Numbers][title]

## Description

Given an integer array, find three numbers whose product is maximum and output the maximum product.
**Example 1:**

```
Input: a = "11", b = "1"
Output: "100"
```

**Example 2:**

```
Input: a = "1010", b = "1011"
Output: "10101"
```

**Tags:** Math, String

## 题意
> 求2数之和

## 题解

### 思路1
> 三个数的最大乘积，必然是三个最大的数的乘积或者两个最小的（负）数与最大数的乘积。
  通过线性扫描或者排序找出最大的三个数和最小的两个数即可。

时间复杂
线性扫描的时间复杂度是 O(n)O(n)；排序的时间复杂度是 O(nlogn)O(nlog⁡n)。

```go
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[n-1]*nums[n-2]*nums[n-3], nums[n-1]*nums[0]*nums[1])
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/maximum-product-of-three-numbers/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
