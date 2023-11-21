# [1814.Count Nice Pairs in an Array][title]

## Description
You are given an array `nums` consists of non-negative integers. Let us define `rev(x)` as the reverse of the non-negative integer `x`. For example, `rev(123) = 321`, and `rev(120) = 21`. A pair of indices (`i, j`) is **nice** if it satisfies all of the following conditions:

- `0 <= i < j < nums.length`
- `nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])`

Return the number of nice pairs of indices. Since that number can be too large, return it **modulo** `10^9 + 7`.

**Example 1:**

```
Input: nums = [42,11,1,97]
Output: 2
Explanation: The two pairs are:
 - (0,3) : 42 + rev(97) = 42 + 79 = 121, 97 + rev(42) = 97 + 24 = 121.
 - (1,2) : 11 + rev(1) = 11 + 1 = 12, 1 + rev(11) = 1 + 11 = 12.
```

**Example 2:**

```
Input: nums = [13,10,35,24,76]
Output: 4
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/count-nice-pairs-in-an-array/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
