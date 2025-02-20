# [942.DI String Match][title]

## Description
A permutation `perm` of `n + 1` integers of all the integers in the range `[0, n]` can be represented as a string `s` of length `n` where:

- `s[i] == 'I'` if `perm[i] < perm[i + 1]`, and
- `s[i] == 'D'` if `perm[i] > perm[i + 1]`.

Given a string `s`, reconstruct the permutation `perm` and return it. If there are multiple valid permutations perm, return **any of them**.

**Example 1:**

```
Input: s = "IDID"
Output: [0,4,1,3,2]
```

**Example 2:**

```
Input: s = "III"
Output: [0,1,2,3]
```

**Example 3:**

```
Input: s = "DDI"
Output: [3,2,0,1]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/di-string-match/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
