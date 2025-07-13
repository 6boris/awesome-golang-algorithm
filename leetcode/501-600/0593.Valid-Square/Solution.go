package Solution

func isVertical(x1, y1, x2, y2 []int) bool {
	// 0
	a := y1[0] - x1[0]
	// -2
	b := y1[1] - x1[1]
	// -2
	c := y2[0] - x2[0]
	// 0
	d := y2[1] - x2[1]
	return a*c+b*d == 0
}

func distance(a, b []int) int {
	x := a[0] - b[0]
	y := a[1] - b[1]
	return x*x + y*y
}
func edgesEqual(x1, y1, x2, y2 []int) bool {
	x1x2 := distance(x1, x2)
	x1y2 := distance(x1, y2)

	if x1x2 != x1y2 {
		return false
	}
	y1x2 := distance(y1, x2)
	y1y2 := distance(y1, y2)
	if y1x2 != y1y2 {
		return false
	}
	if y1x2 != x1y2 {
		return false
	}
	return distance(x1, y1) == distance(x2, y2)
}

func Solution(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	ps := [][]int{p1, p2, p3, p4}
	for i := range 3 {
		for j := i + 1; j < 4; j++ {
			if ps[i][0] == ps[j][0] && ps[i][1] == ps[j][1] {
				return false
			}
		}
	}
	groups := [][][]int{
		{p1, p2, p3, p4},
		{p1, p3, p2, p4},
		{p1, p4, p2, p3},
	}
	for _, g := range groups {
		if !isVertical(g[0], g[1], g[2], g[3]) {
			continue
		}
		if !edgesEqual(g[0], g[1], g[2], g[3]) {
			continue
		}
		return true
	}
	return false
}
