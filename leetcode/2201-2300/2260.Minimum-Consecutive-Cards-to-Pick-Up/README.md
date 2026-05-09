# [2260.Minimum Consecutive Cards to Pick Up][title]

## Description
You are given an integer array `cards` where `cards[i]` represents the **value** of the `ith` card. A pair of cards are **matching** if the cards have the **same** value.

Return the **minimum** number of **consecutive** cards you have to pick up to have a pair of **matching** cards among the picked cards. If it is impossible to have matching cards, return `-1`.

**Example 1:**

```
Input: cards = [3,4,2,3,4,7]
Output: 4
Explanation: We can pick up the cards [3,4,2,3] which contain a matching pair of cards with value 3. Note that picking up the cards [4,2,3,4] is also optimal.
```

**Example 2:**

```
Input: cards = [1,0,5,3]
Output: -1
Explanation: There is no way to pick up a set of consecutive cards that contain a pair of matching cards.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/minimum-consecutive-cards-to-pick-up/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
