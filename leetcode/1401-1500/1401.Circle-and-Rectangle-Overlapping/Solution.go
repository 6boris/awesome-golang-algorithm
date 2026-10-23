package Solution

func Solution(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	x := xCenter
	if xCenter < x1 {
		x = x1
	} else if xCenter > x2 {
		x = x2
	}

	y := yCenter
	if yCenter < y1 {
		y = y1
	} else if yCenter > y2 {
		y = y2
	}

	dx := xCenter - x
	dy := yCenter - y
	return dx*dx+dy*dy <= radius*radius
}
