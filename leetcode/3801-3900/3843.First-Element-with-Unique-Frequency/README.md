# [3843.First Element with Unique Frequency][title]

## Description
You are given an integer array `nums`.

Return an integer denoting the **first** element (scanning from left to right) in `nums` whose **frequency** is **unique**. That is, no other integer appears the same number of times in `nums`. If there is no such element, return -1.

**Example 1:**

```
Input: nums = [20,10,30,30]

Output: 30

Explanation:

20 appears once.
10 appears once.
30 appears twice.
The frequency of 30 is unique because no other integer appears exactly twice.
```

**Example 2:**

```
Input: nums = [20,20,10,30,30,30]

Output: 20

Explanation:

20 appears twice.
10 appears once.
30 appears 3 times.
The frequency of 20, 10, and 30 are unique. The first element that has unique frequency is 20.
```

**Example 3:**

```
Input: nums = [10,10,20,20]

Output: -1

Explanation:

10 appears twice.
20 appears twice.
No element has a unique frequency.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/first-element-with-unique-frequency/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
