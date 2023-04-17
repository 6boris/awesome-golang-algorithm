# [2390.Removing Stars From a String][title]

## Description
You are given a string `s`, which contains stars `*`.

In one operation, you can:

- Choose a star in `s`.
- Remove the closest **non-star** character to its **left**, as well as remove the star itself.

Return the string after **all** stars have been removed.

**Note:**

- The input will be generated such that the operation is always possible.
- It can be shown that the resulting string will always be unique.


**Example 1:**

```
Input: s = "leet**cod*e"
Output: "lecoe"
Explanation: Performing the removals from left to right:
- The closest character to the 1st star is 't' in "leet**cod*e". s becomes "lee*cod*e".
- The closest character to the 2nd star is 'e' in "lee*cod*e". s becomes "lecod*e".
- The closest character to the 3rd star is 'd' in "lecod*e". s becomes "lecoe".
There are no more stars, so we return "lecoe".
```

**Example 2:**

```
Input: s = "erase*****"
Output: ""
Explanation: The entire string is removed, so we return an empty string.
```


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/removing-stars-from-a-string/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
