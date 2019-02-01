package Solution

import (
	"fmt"
	"strconv"
)

var minSum = 99999

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

//	dp solution
func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			minSum := min(triangle[i+1][j], triangle[i+1][j+1])
			minSum += triangle[i][j]
			triangle[i][j] = minSum
		}
	}
	return triangle[0][0]
}

//	dp solution
func minimumTotal2(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
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
