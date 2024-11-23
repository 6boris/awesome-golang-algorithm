package Solution

func Solution(box [][]byte) [][]byte {
	rows, cols := len(box), len(box[0])
	for i := 0; i < rows; i++ {
		for j := cols - 2; j >= 0; j-- {
			if box[i][j] == '.' || box[i][j] == '*' {
				continue
			}
			next := j + 1
			for ; next < cols; next++ {
				if box[i][next] == '#' || box[i][next] == '*' {
					break
				}
			}
			box[i][j] = '.'
			box[i][next-1] = '#'
		}
	}
	res := make([][]byte, cols)
	for i := range cols {
		res[i] = make([]byte, rows)
		for j := 0; j < rows; j++ {
			res[i][j] = box[rows-1-j][i]
		}
	}

	return res
}
