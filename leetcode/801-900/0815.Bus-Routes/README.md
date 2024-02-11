# [815.Bus Routes][title]

## Description
You are given an array `routes` representing bus routes where `routes[i]` is a bus route that the i<sup>th</sup> bus repeats forever.

- For example, if `routes[0] = [1, 5, 7]`, this means that the 0<sup>th</sup> bus travels in the sequence `1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ...` forever.

You will start at the bus stop `source` (You are not on any bus initially), and you want to go to the bus stop `target`. You can travel between bus stops by buses only.

Return the least number of buses you must take to travel from `source` to `target`. Return `-1` if it is not possible.

**Example 1:**

```
Input: routes = [[1,2,7],[3,6,7]], source = 1, target = 6
Output: 2
Explanation: The best strategy is take the first bus to the bus stop 7, then take the second bus to the bus stop 6.
```

**Example 2:**

```
Input: routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
Output: -1
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/bus-routes/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
