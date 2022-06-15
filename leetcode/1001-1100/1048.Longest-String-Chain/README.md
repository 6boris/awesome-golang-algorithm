# [1048.Longest String Chain][title]

## Description
ou are given an array of `words` where each word consists of lowercase English letters.

**word<sub>A</sub>** is a **predecessor** of **word<sub>B</sub>** if and only if we can insert **exactly one** letter anywhere in **word<sub>A</sub>**  **without changing the order of the other characters** to make it equal to **word<sub>B</sub>**.

- For example, `"abc"` is a **predecessor** of `"abac"`, while `"cba"` is not a **predecessor** of `"bcad"`.

A **word chain** is a sequence of words `[word1, word2, ..., wordk]` with `k >= 1`, where **word<sub>1</sub>** is a **predecessor** of **word<sub>2</sub>**, **word<sub>2</sub>** is a **predecessor** of **word<sub>3</sub>**, and so on. A single word is trivially a **word chain** with `k == 1`.

Return the **length** of the **longest possible word chain** with words chosen from the given list of `words`.

**Example 1:**

```
Input: words = ["a","b","ba","bca","bda","bdca"]
Output: 4
Explanation: One of the longest word chains is ["a","ba","bda","bdca"].
```

**Example 2:**

```
Input: words = ["xbc","pcxbcf","xb","cxbc","pcxbc"]
Output: 5
Explanation: All the words can be put in a word chain ["xb", "xbc", "cxbc", "pcxbc", "pcxbcf"].
```

**Example 3:**

```
Input: words = ["abcd","dbqca"]
Output: 1
Explanation: The trivial word chain ["abcd"] is one of the longest word chains.
["abcd","dbqca"] is not a valid word chain because the ordering of the letters is changed.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/longest-string-chain/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
