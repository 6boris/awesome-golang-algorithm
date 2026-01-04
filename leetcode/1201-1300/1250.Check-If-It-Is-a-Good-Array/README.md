# [1250.Check If It Is a Good Array][title]

## Description
Given an array `nums` of positive integers. Your task is to select some subset of `nums`, multiply each element by an integer and add all these numbers. The array is said to be **good** if you can obtain a sum of `1` from the array by any possible subset and multiplicand.

Return `True` if the array is **good** otherwise return `False`.

**Example 1:**

```
Input: nums = [12,5,7,23]
Output: true
Explanation: Pick numbers 5 and 7.
5*3 + 7*(-2) = 1
```

**Example 2:**

```
Input: nums = [29,6,10]
Output: true
Explanation: Pick numbers 29, 6 and 10.
29*1 + 6*(-3) + 10*(-1) = 1
```

**Example 3:**

```
Input: nums = [3,6]
Output: false
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/check-if-it-is-a-good-array/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
