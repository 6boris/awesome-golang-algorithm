# [892.Surface Area of 3D Shapes][title]

## Description
You are given an `n x n` grid where you have placed some `1 x 1 x 1` cubes. Each value `v = grid[i][j]` represents a tower of `v` cubes placed on top of cell `(i, j)`.

After placing these cubes, you have decided to glue any directly adjacent cubes to each other, forming several irregular 3D shapes.

Return the total surface area of the resulting shapes.

**Note**: The bottom face of each shape counts toward its surface area.

**Example 1:**  
![tmp-grid2](./tmp-grid2.jpg)

```
Input: grid = [[1,2],[3,4]]
Output: 34
```

**Example 2:**  

![tmp-grid4](./tmp-grid4.jpg)

```
Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
Output: 32
```

**Example 3:**  

![tmp-grid5](./tmp-grid5.jpg)

```
Input: grid = [[2,2,2],[2,1,2],[2,2,2]]
Output: 46
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/surface-area-of-3d-shapes/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
