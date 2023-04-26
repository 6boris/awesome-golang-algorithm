package Solution

func minAndMax(a, b, c, d int) (int, int) {
	min, max := a, a
	if b > max {
		max = b
	}
	if b < min {
		min = b
	}
	if c > max {
		max = c
	}
	if c < min {
		min = c
	}
	if d > max {
		max = d
	}
	if d < min {
		min = d
	}
	return max, min
}

func Solution(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	// 4 * 4
	width1 := ax2 - ax1
	heigh1 := ay2 - ay1

	width2 := bx2 - bx1
	height2 := by2 - by1
	total := width1*heigh1 + width2*height2

	if bx2 <= ax1 || bx1 >= ax2 || by1 >= ay2 || by2 <= ay1 {
		return total
	}
	topY, downY := minAndMax(ay1, ay2, by1, by2)
	topX, downX := minAndMax(ax1, ax2, bx1, bx2)
	expectWidth := width1 + width2
	realWidth := topX - downX
	expectHeight := heigh1 + height2
	realHeight := topY - downY
	return total - (expectWidth-realWidth)*(expectHeight-realHeight)
}
