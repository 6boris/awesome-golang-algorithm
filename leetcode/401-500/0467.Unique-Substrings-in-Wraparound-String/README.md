# [467.Unique Substrings in Wraparound String][title]

## Description
We define the string `base` to be the infinite wraparound string of `"abcdefghijklmnopqrstuvwxyz"`, so `base` will look like this:

- `"...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd...."`.

Given a string `s`, return the number of **unique non-empty substrings** of `s` are present in `base`.

**Example 1:**

```
Input: s = "a"
Output: 1
Explanation: Only the substring "a" of s is in base.
```

**Example 2:**

```
Input: s = "cac"
Output: 2
Explanation: There are two substrings ("a", "c") of s in base.
```

**Example 3:**

```
Input: s = "zab"
Output: 6
Explanation: There are six substrings ("z", "a", "b", "za", "ab", and "zab") of s in base.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/unique-substrings-in-wraparound-string/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
