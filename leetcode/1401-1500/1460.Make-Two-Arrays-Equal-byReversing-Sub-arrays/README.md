# [1460.Make Two Arrays Equal by Reversing Sub-arrays][title]

## Description
You are given two integer arrays of equal length `target` and `arr`. In one step, you can select any `non-empty sub-array` of `arr` and reverse it. You are allowed to make any number of steps.

Return `true` if you can make `arr` equal to `target` or `false` otherwise.

**Example 1:**

```
Input: target = [1,2,3,4], arr = [2,4,1,3]
Output: true
Explanation: You can follow the next steps to convert arr to target:
1- Reverse sub-array [2,4,1], arr becomes [1,4,2,3]
2- Reverse sub-array [4,2], arr becomes [1,2,4,3]
3- Reverse sub-array [4,3], arr becomes [1,2,3,4]
There are multiple ways to convert arr to target, this is not the only way to do so.
```

**Example 2:**

```
Input: target = [7], arr = [7]
Output: true
Explanation: arr is equal to target without any reverses.
```

**Example 3:**

```
Input: target = [3,7,9], arr = [3,7,11]
Output: false
Explanation: arr does not have value 9 and it can never be converted to target.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/make-two-arrays-equal-by-reversing-sub-arrays/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
