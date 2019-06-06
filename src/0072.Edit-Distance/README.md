# [72. Edit Distance][title]

## Description

Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

- Insert a character
- Delete a character
- Replace a character

**Example 1:**

```
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation: 
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
```

**Example 2:**

```
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation: 
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
```

**Tags:** Math, String

## 题意
>给定两个单词 word1 和 word2 ，请问最少需要进行多少次操作可以将 word1 变成 word2 。

## 题解

### 思路1
> (动态规划) O(n2)

经典的编辑距离问题。

状态表示：`f[i,j]` 表示将 `word1` 的前 `i` 个字符变成 `word2` 的前 `j` 个字符，最少需要进行多少次操作。
状态转移，一共有四种情况：

- 将`word1[i]` 删除或在`word2[j]`后面添加`word1[i]`，则其操作次数等于`f[i−1,j]+1`；
- 将`word2[j]`删除或在`word1[i]`后面添加`word2[j]`，则其操作次数等于 `f[i,j−1]+1`；
- 如果`word1[i]=word2[j]`，则其操作次数等于 `f[i−1,j−1]`；
- 如果`word1[i]≠word2[j]`，则其操作次数等于 `f[i−1,j−1]+1`；

> 时间复杂度分析：

- 状态数 O(n2)O(n2)，状态转移复杂度是 O(1)O(1)，所以总时间复杂度是 O(n2)O(n2)。

```go

```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
