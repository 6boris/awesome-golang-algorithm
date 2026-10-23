# [3162.Find the Number of Good Pairs I][title]

## Description
You are given 2 integer arrays `nums1` and `nums2` of lengths `n` and `m` respectively. You are also given a **positive** integer `k`.

A pair `(i, j)` is called **good** if `nums1[i]` is divisible by `nums2[j] * k (0 <= i <= n - 1, 0 <= j <= m - 1)`.

Return the total number of **good** pairs.

**Example 1:**

```
Input: nums1 = [1,3,4], nums2 = [1,3,4], k = 1

Output: 5

Explanation:

The 5 good pairs are (0, 0), (1, 0), (1, 1), (2, 0), and (2, 2).
```

**Example 2:**

```
Input: nums1 = [1,2,4,12], nums2 = [2,4], k = 3

Output: 2

Explanation:

The 2 good pairs are (3, 0) and (3, 1).
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/find-the-number-of-good-pairs-i/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
