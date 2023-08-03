package Solution

func Solution(matrix [][]int, target int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	cache := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		cache[i] = make([]int, cols+1)
	}

	/*
		   0 1 0
		   1 1 1
		   0 1 0
			-- [5 4 1 0]
			-- [4 3 1 0]
			-- [1 1 0 0]
			-- [0 0 0 0]
	*/
	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			cache[r][c] = matrix[r][c] + cache[r+1][c] + cache[r][c+1] - cache[r+1][c+1]
		}
	}

	ans := 0

	for x1 := 0; x1 < rows; x1++ {
		for y1 := 0; y1 < cols; y1++ {
			for x2 := x1; x2 < rows; x2++ {
				for y2 := y1; y2 < cols; y2++ {
					diff := cache[x1][y1] + cache[x2+1][y2+1] - cache[x1][y2+1] - cache[x2+1][y1]
					//fmt.Printf("(%d,%d), (%d,%d) = %d\n", x1, y1, x2, y2, diff)
					if diff == target {
						ans++
					}
				}
			}
		}
	}
	return ans
}
