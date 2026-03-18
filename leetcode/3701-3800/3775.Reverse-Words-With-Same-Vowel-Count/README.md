# [3775.Reverse Words With Same Vowel Count][title]

## Description
You are given a string `s` consisting of lowercase English words, each separated by a single space.

Determine how many vowels appear in the **first** word. Then, reverse each following word that has the **same vowel count**. Leave all remaining words unchanged.

Return the resulting string.

Vowels are `'a'`, `'e'`, `'i'`, `'o'`, and `'u'`.

**Example 1:**

```
Input: s = "cat and mice"

Output: "cat dna mice"

Explanation:​​​​​​​

The first word "cat" has 1 vowel.
"and" has 1 vowel, so it is reversed to form "dna".
"mice" has 2 vowels, so it remains unchanged.
Thus, the resulting string is "cat dna mice".
```

**Example 2:**

```
Input: s = "book is nice"

Output: "book is ecin"

Explanation:

The first word "book" has 2 vowels.
"is" has 1 vowel, so it remains unchanged.
"nice" has 2 vowels, so it is reversed to form "ecin".
Thus, the resulting string is "book is ecin".
```

**Example 3:**

```
Input: s = "banana healthy"

Output: "banana healthy"

Explanation:

The first word "banana" has 3 vowels.
"healthy" has 2 vowels, so it remains unchanged.
Thus, the resulting string is "banana healthy".
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/reverse-words-with-same-vowel-count/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
