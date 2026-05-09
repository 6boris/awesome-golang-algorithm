# [3661.Maximum Walls Destroyed by Robots][title]

## Description
There is an endless straight line populated with some `robots` and `walls`. You are given integer arrays robots, distance, and walls:

- `robots[i]` is the position of the `ith` robot.
- `distance[i]` is the **maximum** distance the `ith` robot's bullet can travel.
- `walls[j]` is the position of the `jth` wall.

Every robot has **one** bullet that can either fire to the left or the right **at most** `distance[i]` meters.

A bullet destroys every wall in its path that lies within its range. Robots are fixed obstacles: if a bullet hits another robot before reaching a wall, it **immediately stops** at that robot and cannot continue.

Return the **maximum** number of **unique** walls that can be destroyed by the robots.

Notes:

- A wall and a robot may share the same position; the wall can be destroyed by the robot at that position.
- Robots are not destroyed by bullets.

**Example 1:**

```
Input: robots = [4], distance = [3], walls = [1,10]

Output: 1

Explanation:

robots[0] = 4 fires left with distance[0] = 3, covering [1, 4] and destroys walls[0] = 1.
Thus, the answer is 1.
```

**Example 2:**

```
Input: robots = [10,2], distance = [5,1], walls = [5,2,7]

Output: 3

Explanation:

robots[0] = 10 fires left with distance[0] = 5, covering [5, 10] and destroys walls[0] = 5 and walls[2] = 7.
robots[1] = 2 fires left with distance[1] = 1, covering [1, 2] and destroys walls[1] = 2.
Thus, the answer is 3.
```

**Example 3:**

```
Input: robots = [1,2], distance = [100,1], walls = [10]

Output: 0

Explanation:

In this example, only robots[0] can reach the wall, but its shot to the right is blocked by robots[1]; thus the answer is 0.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/maximum-walls-destroyed-by-robots/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
