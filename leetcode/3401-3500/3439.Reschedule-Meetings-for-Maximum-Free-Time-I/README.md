# [3439.Reschedule Meetings for Maximum Free Time I][title]

## Description
You are given an integer `eventTime` denoting the duration of an event, where the event occurs from time `t = 0` to time `t = eventTime`.

You are also given two integer arrays `startTime` and `endTime`, each of length `n`. These represent the start and end time of `n` **non-overlapping** meetings, where the `ith` meeting occurs during the time `[startTime[i], endTime[i]]`.

You can reschedule **at most** `k` meetings by moving their start time while maintaining the **same duration**, to **maximize** the **longest** continuous period of free time during the event.

The **relative** order of all the meetings should stay the same and they should remain non-overlapping.

Return the **maximum** amount of free time possible after rearranging the meetings.

**Note** that the meetings can **not** be rescheduled to a time outside the event.

**Example 1:**  

![1](./1.png)

```
Input: eventTime = 5, k = 1, startTime = [1,3], endTime = [2,5]

Output: 2

Explanation:

Reschedule the meeting at [1, 2] to [2, 3], leaving no meetings during the time [0, 2].
```

**Example 2:**  

![2](./2.png)

```
Input: eventTime = 10, k = 1, startTime = [0,2,9], endTime = [1,4,10]

Output: 6

Explanation:

Reschedule the meeting at [2, 4] to [1, 3], leaving no meetings during the time [3, 9].
```

**Example 3:**

```
Input: eventTime = 5, k = 2, startTime = [0,1,2,3,4], endTime = [1,2,3,4,5]

Output: 0

Explanation:

There is no time during the event not occupied by meetings
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-i/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
