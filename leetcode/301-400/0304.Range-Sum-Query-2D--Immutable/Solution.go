package Solution

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	for row := 0; row < len(matrix); row++ {
		for col := 1; col < len(matrix[row]); col++ {
			matrix[row][col] += matrix[row][col-1]
		}
	}

	return NumMatrix{matrix: matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := 0
	for row := row1; row <= row2; row++ {
		ans += this.matrix[row][col2]
		if col1 > 0 {
			ans -= this.matrix[row][col1-1]
		}
	}
	return ans
}

func Solution(matrix [][]int, coordinates [][4]int) []int {
	o := Constructor(matrix)
	ans := make([]int, 0)
	for _, coordinate := range coordinates {
		ans = append(ans, o.SumRegion(coordinate[0], coordinate[1], coordinate[2], coordinate[3]))
	}

	return ans
}
