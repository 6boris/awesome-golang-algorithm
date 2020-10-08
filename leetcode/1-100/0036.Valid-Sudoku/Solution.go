package Solution

func isValidSudoku(board [][]byte) bool {
	n := len(board)
	rectSets := make([]int16, n*n/9) // sets for inner rectangles
	for i := 0; i < n; i++ {
		rowSet, colSet := int16(0), int16(0) // sets for a row and a column
		for j := 0; j < n; j++ {
			// checking a column and an inner rectangle
			if num := board[i][j]; num != '.' {
				numBit := int16(1 << (num - '0')) // calc a #num bit
				if rowSet&numBit > 0 {
					return false
				}
				rowSet |= numBit // set the bit

				// set the bit for a inner rectangle: ((n/3)*(i/3), j/3) - its coordinates
				if rectSets[(n/3)*(i/3)+j/3]&numBit > 0 {
					return false
				}
				rectSets[(n/3)*(i/3)+j/3] |= numBit
			}

			// checking column
			if num := board[j][i]; num != '.' {
				numBit := int16(1 << (num - '0'))
				if colSet&numBit > 0 {
					return false
				}
				colSet ^= numBit
			}
		}
	}
	return true
}
