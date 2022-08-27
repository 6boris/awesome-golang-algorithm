# [1736.Latest Time by Replacing Hidden Digits][title]

## Description
You are given a string `time` in the form of `hh:mm`, where some of the digits in the string are hidden (represented by `?`).

The valid times are those inclusively between `00:00` and `23:59`.

Return the latest valid time you can get from `time` by replacing the hidden digits.

**Example 1:**

```
Input: time = "2?:?0"
Output: "23:50"
Explanation: The latest hour beginning with the digit '2' is 23 and the latest minute ending with the digit '0' is 50.
```

**Example 2:**

```
Input: time = "0?:3?"
Output: "09:39"
```

**Example 3:**

```
Input: time = "1?:22"
Output: "19:22"
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/latest-time-by-replacing-hidden-digits/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
