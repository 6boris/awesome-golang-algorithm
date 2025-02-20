# [3151.Special Array I][title]

## Description
An array is considered **special** if every pair of its adjacent elements contains two numbers with different parity.

You are given an array of integers `nums`. Return `true` if `nums` is a **special** array, otherwise, return `false`.

**Example 1:**

```
Input: nums = [1]

Output: true

Explanation:

There is only one element. So the answer is true.
```

**Example 2:**

```
Input: nums = [2,1,4]

Output: true

Explanation:

There is only two pairs: (2,1) and (1,4), and both of them contain numbers with different parity. So the answer is true.
```

**Example 3:**

```
Input: nums = [4,3,1,6]

Output: false

Explanation:

nums[1] and nums[2] are both odd. So the answer is false.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/special-array-i/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
