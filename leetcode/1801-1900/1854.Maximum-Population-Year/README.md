# [1854.Maximum Population Year][title]

## Description
You are given a 2D integer array `logs` where each `logs[i] = [birthi, deathi]` indicates the birth and death years of the i<sup>th</sup> person.

The **population** of some year `x` is the number of people alive during that year. The i<sup>th</sup> person is counted in year `x`'s population if `x` is in the **inclusive** range `[birthi, deathi - 1]`. Note that the person is **not** counted in the year that they die.

Return the **earliest** year with the **maximum population**.


**Example 1:**

```
Input: logs = [[1993,1999],[2000,2010]]
Output: 1993
Explanation: The maximum population is 1, and 1993 is the earliest year with this population.
```

**Example 2:**

```
Input: logs = [[1950,1961],[1960,1971],[1970,1981]]
Output: 1960
Explanation: 
The maximum population is 2, and it had happened in years 1960 and 1970.
The earlier year between them is 1960.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/maximum-population-year/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
