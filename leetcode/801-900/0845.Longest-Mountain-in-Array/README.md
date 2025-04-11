# [845.Longest Mountain in Array][title]

## Description
You may recall that an array `arr` is a **mountain array** if and only if:

- `arr.length >= 3`
- There exists some index `i` **(0-indexed)** with `0 < i < arr.length - 1` such that:

    - `arr[0] < arr[1] < ... < arr[i - 1] < arr[i]`
    - `arr[i] > arr[i + 1] > ... > arr[arr.length - 1]`

Given an integer array `arr`, return the length of the longest subarray, which is a mountain. Return `0` if there is no mountain subarray.

**Example 1:**

```
Input: arr = [2,1,4,7,3,2,5]
Output: 5
Explanation: The largest mountain is [1,4,7,3,2] which has length 5.
```

**Example 2:**

```
Input: arr = [2,2,2]
Output: 0
Explanation: There is no mountain.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/longest-mountain-in-array/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
