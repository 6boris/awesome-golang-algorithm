# [2053.Kth Distinct String in an Array][title]

## Description
A **distinct string** is a string that is present only **once** in an array.

Given an array of strings `arr`, and an integer `k`, return the `kth` **distinct string** present in `arr`. If there are **fewer** than `k` distinct strings, return an **empty string** `""`.

Note that the strings are considered in the **order in which they appear** in the array.

**Example 1:**

```
Input: arr = ["d","b","c","b","c","a"], k = 2
Output: "a"
Explanation:
The only distinct strings in arr are "d" and "a".
"d" appears 1st, so it is the 1st distinct string.
"a" appears 2nd, so it is the 2nd distinct string.
Since k == 2, "a" is returned.
```

**Example 2:**

```
Input: arr = ["aaa","aa","a"], k = 1
Output: "aaa"
Explanation:
All strings in arr are distinct, so the 1st string "aaa" is returned.
```

**Example 3:**

```
Input: arr = ["a","b","a"], k = 3
Output: ""
Explanation:
The only distinct string is "b". Since there are fewer than 3 distinct strings, we return an empty string "".
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/kth-distinct-string-in-an-array/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
