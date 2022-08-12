# [858.Mirror Reflection][title]

## Description
There is a special square room with mirrors on each of the four walls. Except for the southwest corner, there are receptors on each of the remaining corners, numbered `0`, `1`, and `2`.

The square room has walls of length `p` and a laser ray from the southwest corner first meets the east wall at a distance q from the 0<sup>th</sup> receptor.

Given the two integers `p` and `q`, return the number of the receptor that the ray meets first.

The test cases are guaranteed so that the ray will meet a receptor eventually.


**Example 1:**  

![reflection](./reflection.png)

```
Input: p = 2, q = 1
Output: 2
Explanation: The ray meets receptor 2 the first time it gets reflected back to the left wall.
```

**Example 2:**

```
Input: p = 3, q = 1
Output: 1
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/mirror-reflection/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
