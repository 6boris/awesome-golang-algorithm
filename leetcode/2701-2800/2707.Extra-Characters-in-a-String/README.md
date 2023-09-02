# [2707.Extra Characters in a String][title]

## Description
You are given a **0-indexed** string s and a `dictionary` of words dictionary. You have to break s into one or more **non-overlapping** substrings such that each substring is present in `dictionary`. There may be some **extra characters** in s which are not present in any of the substrings.

Return the minimum number of extra characters left over if you break up `s` optimally.

**Example 1:**

```
Input: s = "leetscode", dictionary = ["leet","code","leetcode"]
Output: 1
Explanation: We can break s in two substrings: "leet" from index 0 to 3 and "code" from index 5 to 8. There is only 1 unused character (at index 4), so we return 1.
```

**Example 2:**

```
Input: s = "sayhelloworld", dictionary = ["hello","world"]
Output: 3
Explanation: We can break s in two substrings: "hello" from index 3 to 7 and "world" from index 8 to 12. The characters at indices 0, 1, 2 are not used in any substring and thus are considered as extra characters. Hence, we return 3.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/extra-characters-in-a-string
[me]: https://github.com/kylesliu/awesome-golang-algorithm
