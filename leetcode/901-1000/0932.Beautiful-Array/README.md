# [932.Beautiful Array][title]

## Description
An array `nums` of length `n` is **beautiful** if:

- `nums` is a permutation of the integers in the range `[1, n]`.
- For every `0 <= i < j < n`, there is no index `k` with `i < k < j` where `2 * nums[k] == nums[i] + nums[j]`.

Given the integer `n`, return any **beautiful** array `nums` of length `n`. There will be at least one valid answer for the given `n`.

**Example 1:**

```
Input: n = 4
Output: [2,1,4,3]
```

**Example 2:**

```
Input: n = 5
Output: [3,1,2,5,4]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/beautiful-array/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
