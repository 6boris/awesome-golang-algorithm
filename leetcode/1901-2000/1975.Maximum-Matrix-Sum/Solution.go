package Solution

func Solution(matrix [][]int) int64 {
	ans, neg := int64(0), 0
	_min := int64(100001)
	zero := 0
	rows, cols := len(matrix), len(matrix[0])
	for i := range rows {
		for j := range cols {
			cur := int64(matrix[i][j])
			if cur == 0 {
				zero++
			}
			if cur < 0 {
				neg++
				cur = -cur
			}
			ans += cur
			_min = min(_min, cur)
		}
	}
	if neg&1 == 1 && zero == 0 {
		ans -= 2 * _min
	}
	return ans
}
