# [1957.Delete Characters to Make Fancy String][title]

## Description
A **fancy string** is a string where no **three consecutive** characters are equal.

Given a string `s`, delete the **minimum** possible number of characters from `s` to make it `fancy`.

Return the final string after the deletion. It can be shown that the answer will always be `unique`.

**Example 1:**

```
Input: s = "leeetcode"
Output: "leetcode"
Explanation:
Remove an 'e' from the first group of 'e's to create "leetcode".
No three consecutive characters are equal, so return "leetcode".
```

**Example 2:**

```
Input: s = "aaabaaaa"
Output: "aabaa"
Explanation:
Remove an 'a' from the first group of 'a's to create "aabaaaa".
Remove two 'a's from the second group of 'a's to create "aabaa".
No three consecutive characters are equal, so return "aabaa".
```

**Example 3:**

```
Input: s = "aab"
Output: "aab"
Explanation: No three consecutive characters are equal, so return "aab".
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/delete-characters-to-make-fancy-string/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
