# [980. Unique Paths III][title]

## Description

On a 2-dimensional grid, there are 4 types of squares:

1 represents the starting square.  There is exactly one starting square.
2 represents the ending square.  There is exactly one ending square.
0 represents empty squares we can walk over.
-1 represents obstacles that we cannot walk over.
Return the number of 4-directional walks from the starting square to the ending square, that walk over every non-obstacle square exactly once.

**Example 1:**

```
Input: [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
Output: 2
Explanation: We have the following two paths: 
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
```

**Example 2:**

```
Input: [[1,0,0,0],[0,0,0,0],[0,0,0,2]]
Output: 4
Explanation: We have the following four paths: 
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
```
**Example 3:**

```
Input: [[0,1],[2,0]]
Output: 0
Explanation: 
There is no path that walks over every empty square exactly once.
Note that the starting and ending square can be anywhere in the grid.4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
```


**Tags:** Math, String

## 题意
>求路径比probelem 63 要复杂许多

## 题解

### 思路1
> DFS
```go
//	DFS Brute Force
func uniquePathsIII(grid [][]int) int {
	sx, sy, n := -1, -1, 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				n++
			} else if grid[i][j] == 1 {
				sx, sy = j, i
			}

		}
	}
	return dfs(grid, sx, sy, n)

}

func dfs(grid [][]int, x, y, n int) int {
	//	terminator
	if x < 0 || x == len(grid[0]) ||
		y < 0 || y == len(grid) ||
		grid[y][x] == -1 {
		return 0
	}
	if grid[y][x] == 2 {
		if n == 0 {
			return 1
		} else {
			return 0
		}
	}

	//	process
	grid[y][x] = -1

	//	drill down
	path := dfs(grid, x+1, y, n-1) +
		dfs(grid, x-1, y, n-1) +
		dfs(grid, x, y+1, n-1) +
		dfs(grid, x, y-1, n-1)

	//	clear data
	grid[y][x] = 0

	return path
}

```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
