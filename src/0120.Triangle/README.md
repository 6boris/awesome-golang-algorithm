# [120. Triangle][title]

## Description
Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
**Example 1:**

```
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
```
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

**Tags:** Math, String

## 题意
>求三角形的最短路径.

## 题解

### 思路1
> DFS

```go
//	dfs solution
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	_dfs(triangle, 0, 0, "", 0)

	return minSum
}

func _dfs(triangle [][]int, i, j int, path string, sum int) int {
	//	terminator
	if i == len(triangle)-1 {
		sum += triangle[i][j]
		path += strconv.Itoa(triangle[i][j]) + " #"
		fmt.Println(path + "with sum: " + strconv.Itoa(sum))

		if sum < minSum {
			minSum = sum
			fmt.Println(minSum)
		}
		return sum
	}

	//	process
	sum += triangle[i][j]
	path += strconv.Itoa(triangle[i][j]) + " -> "

	//	drill down
	_dfs(triangle, i+1, j, path, sum)
	_dfs(triangle, i+1, j+1, path, sum)

	//	clear states
	return 0
}
```

### 思路2
> DP
```go
//	dp solution
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			minSum := min(triangle[i+1][j],triangle[i+1][j+1])
			minSum += triangle[i][j]
			triangle[i][j] = minSum
		}
	}
	return triangle[0][0]
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/triangle/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
