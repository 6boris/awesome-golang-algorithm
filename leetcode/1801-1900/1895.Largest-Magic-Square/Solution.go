package Solution

type Cell1895 struct {
	rowSum, colSum int
}

func Solution(grid [][]int) int {
	ret, edgeLen, target := 1, 0, 0
	topLeft, bottomLeft, lastIndex := 0, 0, 0
	n, m := len(grid), len(grid[0])
	sum := make([][]Cell1895, n+1)
	for i := range n + 1 {
		sum[i] = make([]Cell1895, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum[i+1][j+1].rowSum = sum[i+1][j].rowSum + grid[i][j]
			sum[i+1][j+1].colSum = sum[i][j+1].colSum + grid[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			edgeLen = min(m-j, n-i)
			for el := 2; el <= edgeLen; el++ {
				topLeft = grid[i][j]
				lastIndex = i + el - 1
				bottomLeft = grid[lastIndex][j]

				target = sum[i+el][j+1].colSum - sum[i][j+1].colSum
				row := sum[i+1][j+el].rowSum - sum[i+1][j].rowSum
				if row != target {
					continue
				}
				k := 1
				for ; k < el; k++ {
					topLeft += grid[i+k][j+k]
					bottomLeft += grid[lastIndex-k][j+k]
					one := sum[i+el][j+1+k].colSum - sum[i][j+1+k].colSum
					two := sum[i+1+k][j+el].rowSum - sum[i+1+k][j].rowSum
					if one != target || two != target {
						break
					}
				}
				if k == el && topLeft == target && bottomLeft == target {
					ret = max(ret, el)
				}
			}
		}
	}
	return ret
}
