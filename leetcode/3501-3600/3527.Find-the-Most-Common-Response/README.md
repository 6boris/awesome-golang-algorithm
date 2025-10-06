# [3527.Find the Most Common Response][title]

## Description
You are given a 2D string array `responses` where each `responses[i]` is an array of strings representing survey responses from the `ith` day.

Return the **most common** response across all days after removing **duplicate** responses within each `responses[i]`. If there is a tie, return the lexicographically smallest response.

**Example 1:**

```
Input: responses = [["good","ok","good","ok"],["ok","bad","good","ok","ok"],["good"],["bad"]]

Output: "good"

Explanation:

After removing duplicates within each list, responses = [["good", "ok"], ["ok", "bad", "good"], ["good"], ["bad"]].
"good" appears 3 times, "ok" appears 2 times, and "bad" appears 2 times.
Return "good" because it has the highest frequency.
```

**Example 2:**

```
Input: responses = [["good","ok","good"],["ok","bad"],["bad","notsure"],["great","good"]]

Output: "bad"

Explanation:

After removing duplicates within each list we have responses = [["good", "ok"], ["ok", "bad"], ["bad", "notsure"], ["great", "good"]].
"bad", "good", and "ok" each occur 2 times.
The output is "bad" because it is the lexicographically smallest amongst the words with the highest frequency.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/find-the-most-common-response/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
