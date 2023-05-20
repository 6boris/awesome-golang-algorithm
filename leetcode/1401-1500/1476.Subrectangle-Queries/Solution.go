package Solution

type op struct {
	name                   string
	row1, col1, row2, col2 int
	v                      int
}
type SubrectangleQueries struct {
	rec [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rec: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for r := row1; r <= row2; r++ {
		for c := col1; c <= col2; c++ {
			this.rec[r][c] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rec[row][col]
}

func Solution(rectangle [][]int, ops []op) []int {
	c := Constructor(rectangle)
	ans := make([]int, 0)
	for _, o := range ops {
		if o.name == "UpdateSubrectangle" {
			c.UpdateSubrectangle(o.row1, o.col1, o.row2, o.col2, o.v)
			continue
		}
		ans = append(ans, c.GetValue(o.row1, o.col1))
	}
	return ans
}
