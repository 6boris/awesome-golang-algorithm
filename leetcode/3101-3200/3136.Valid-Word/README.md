# [3136.Valid Word][title]

## Description
A word is considered **valid** if:

- It contains a **minimum** of 3 characters.
- It contains only digits (0-9), and English letters (uppercase and lowercase).
- It includes **at least** one **vowel**.
- It includes **at least** one **consonant**.

You are given a string `word`.

Return `true` if `word` is valid, otherwise, return `false`.

**Notes**:

- `'a'`, `'e'`, `'i'`, `'o'`, `'u'`, and their uppercases are **vowels**.
- A **consonant** is an English letter that is not a vowel.


**Example 1:**

```
Input: word = "234Adas"

Output: true

Explanation:

This word satisfies the conditions.
```

**Example 2:**

```
Input: word = "b3"

Output: false

Explanation:

The length of this word is fewer than 3, and does not have a vowel.
```

**Example 3:**

```
Input: word = "a3$e"

Output: false

Explanation:

This word contains a '$' character and does not have a consonant.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/valid-word/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
