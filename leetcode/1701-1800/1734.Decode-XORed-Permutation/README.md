# [1734.Decode XORed Permutation][title]

## Description
There is an integer array `perm` that is a permutation of the first `n` positive integers, where `n` is always `odd`.

It was encoded into another integer array `encoded` of length `n - 1`, such that `encoded[i] = perm[i] XOR perm[i + 1]`. For example, if `perm = [1,3,2]`, then `encoded = [2,1]`.

Given the `encoded` array, return the original array `perm`. It is guaranteed that the answer exists and is unique.

**Example 1:**

```
Input: encoded = [3,1]
Output: [1,2,3]
Explanation: If perm = [1,2,3], then encoded = [1 XOR 2,2 XOR 3] = [3,1]
```

**Example 2:**

```
Input: encoded = [6,5,4,6]
Output: [2,4,1,5,3]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/decode-xored-permutation/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
