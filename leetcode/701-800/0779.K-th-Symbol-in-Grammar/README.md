# [779.K-th Symbol in Grammar][title]

## Description
We build a table of `n` rows (**1-indexed**). We start by writing `0` in the 1<sup>st</sup> row. Now in every subsequent row, we look at the previous row and replace each occurrence of `0` with `01`, and each occurrence of `1` with `10`.

- For example, for `n = 3`, the 1<sup>st</sup> row is `0`, the 2<sup>nd</sup> row is `01`, and the 3<sup>rd</sup> row is `0110`.

Given two integer `n` and `k`, return the k<sup>th<sup> (**1-indexed**) symbol in the n<sup>th</sup> row of a table of `n` rows.

**Example 1:**

```
Input: n = 1, k = 1
Output: 0
Explanation: row 1: 0
```

**Example 2:**

```
Input: n = 2, k = 1
Output: 0
Explanation: 
row 1: 0
row 2: 01
```

**Example 3:**

```
Input: n = 2, k = 2
Output: 1
Explanation: 
row 1: 0
row 2: 01
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/k-th-symbol-in-grammar/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
