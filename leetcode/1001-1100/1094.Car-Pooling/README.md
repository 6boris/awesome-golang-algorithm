# [1094.Car Pooling][title]

## Description
There is a car with `capacity` empty seats. The vehicle only drives east (i.e., it cannot turn around and drive west).

You are given the integer `capacity` and an array `trips` where trip[i] = [numPassengers<sub>i</sub>, from<sub>i</sub>, to<sub>i</sub>] indicates that the i<sup>th</sup> trip has numPassengers<sub>i</sub> passengers and the locations to pick them up and drop them off are from<sub>i</sub> and to<sub>i</sub> respectively. The locations are given as the number of kilometers due east from the car's initial location.

Return `true` if it is possible to pick up and drop off all passengers for all the given trips, or `false` otherwise.  

**Example 1:**

```
Input: trips = [[2,1,5],[3,3,7]], capacity = 4
Output: false
```

**Example 2:**

```
Input: trips = [[2,1,5],[3,3,7]], capacity = 5
Output: true
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/car-pooling/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
