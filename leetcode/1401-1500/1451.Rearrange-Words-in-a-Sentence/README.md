# [1451.Rearrange Words in a Sentence][title]

## Description
Given a sentence `text` (A sentence is a string of space-separated words) in the following format:

- First letter is in upper case.
- Each word in `text` are separated by a single space.

Your task is to rearrange the words in text such that all words are rearranged in an increasing order of their lengths. If two words have the same length, arrange them in their original order.

Return the new text following the format shown above.

**Example 1:**

```
Input: text = "Leetcode is cool"
Output: "Is cool leetcode"
Explanation: There are 3 words, "Leetcode" of length 8, "is" of length 2 and "cool" of length 4.
Output is ordered by length and the new first word starts with capital letter.
```

**Example 2:**

```
Input: text = "Keep calm and code on"
Output: "On and keep calm code"
Explanation: Output is ordered as follows:
"On" 2 letters.
"and" 3 letters.
"keep" 4 letters in case of tie order by position in original text.
"calm" 4 letters.
"code" 4 letters.
```

**Example 3:**

```
Input: text = "To be or not to be"
Output: "To be or to be not"
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/rearrange-words-in-a-sentence
[me]: https://github.com/kylesliu/awesome-golang-algorithm
