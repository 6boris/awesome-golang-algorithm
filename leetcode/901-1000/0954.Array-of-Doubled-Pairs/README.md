# [954.Array of Doubled Pairs][title]

## Description
Given an integer array of even length `arr`, return `true` if it is possible to reorder arr such that `arr[2 * i + 1] = 2 * arr[2 * i]` for every `0 <= i < len(arr) / 2`, or `false` otherwise.

**Example 1:**

```
Input: arr = [3,1,3,6]
Output: false
```

**Example 2:**

```
Input: arr = [2,1,2,6]
Output: false
```

**Example 3:**

```
Input: arr = [4,-2,2,-4]
Output: true
Explanation: We can take two groups, [-2,-4] and [2,4] to form [-2,-4,2,4] or [2,4,-2,-4].
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/array-of-doubled-pairs/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
