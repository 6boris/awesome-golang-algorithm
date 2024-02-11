# [1897.Redistribute Characters to Make All Strings Equal][title]

## Description
You are given an array of strings `words` (0-indexed).

In one operation, pick two **distinct** indices `i` and `j`, where `words[i]` is a non-empty string, and move **any** character from `words[i]` to any position in `words[j]`.

Return `true` if you can make **every** string in `wrods` **equal** using **any** number of operations, and `false` otherwise.

**Example 1:**

```
Input: words = ["abc","aabc","bc"]
Output: true
Explanation: Move the first 'a' in words[1] to the front of words[2],
to make words[1] = "abc" and words[2] = "abc".
All the strings are now equal to "abc", so return true.
```

**Example 2:**

```
Input: words = ["ab","a"]
Output: false
Explanation: It is impossible to make all the strings equal using the operation.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
