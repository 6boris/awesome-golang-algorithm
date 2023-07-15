# [1751.Maximum Number of Events That Can Be Attended II][title]

## Description
You are given an array of `events` where `events[i] = [startDayi, endDayi, valuei]`. The i<sup>th</sup> event starts at startDay<sub>i</sub> and ends at endDay<sub>i</sub>, and if you attend this event, you will receive a value of value<sub>i</sub>. You are also given an integer `k` which represents the maximum number of events you can attend.

You can only attend one event at a time. If you choose to attend an event, you must attend the **entire** event. Note that the end day is **inclusive**: that is, you cannot attend two events where one of them starts and the other ends on the same day.

Return the **maximum sum** of values that you can receive by attending events.

**Example 1:**  

![example1](./screenshot-2021-01-11-at-60048-pm.png)

```
Input: events = [[1,2,4],[3,4,3],[2,3,1]], k = 2
Output: 7
Explanation: Choose the green events, 0 and 1 (0-indexed) for a total value of 4 + 3 = 7.
```


**Example 2:**  

![example2](./screenshot-2021-01-11-at-60150-pm.png)

```
Input: events = [[1,2,4],[3,4,3],[2,3,10]], k = 2
Output: 10
Explanation: Choose event 2 for a total value of 10.
Notice that you cannot attend any other event as they overlap, and that you do not have to attend k events.
```

**Example 3:**  

![example3](./screenshot-2021-01-11-at-60703-pm.png)

```
Input: events = [[1,1,1],[2,2,2],[3,3,3],[4,4,4]], k = 3
Output: 9
Explanation: Although the events do not overlap, you can only attend 3 events. Pick the highest valued three.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended-ii/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
