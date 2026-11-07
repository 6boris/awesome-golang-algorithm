# [3582.Generate Tag for Video Caption][title]

## Description
You are given a string `caption` representing the caption for a video.

The following actions must be performed **in order** to generate a **valid tag** for the video:

1. **Combine all words** in the string into a single camelCase string prefixed with `'#'`. A camelCase string is one where the first letter of all words except the first one is capitalized. All characters after the first character in **each** word must be lowercase.

2. **Remove** all characters that are not an English letter, **except** the first `'#'`.

3. **Truncate** the result to a maximum of 100 characters.

Return the **tag** after performing the actions on `caption`.

**Example 1:**

```
Input: caption = "Leetcode daily streak achieved"

Output: "#leetcodeDailyStreakAchieved"

Explanation:

The first letter for all words except "leetcode" should be capitalized.
```

**Example 2:**

```
Input: caption = "can I Go There"

Output: "#canIGoThere"

Explanation:

The first letter for all words except "can" should be capitalized.
```

**Example 3:**

```
Input: caption = "hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"

Output: "#hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"

Explanation:

Since the first word has length 101, we need to truncate the last two letters from the word.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/generate-tag-for-video-caption/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
