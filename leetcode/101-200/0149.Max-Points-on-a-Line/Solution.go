package Solution

import "fmt"

func Solution(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	_max := 0
	for idx := 1; idx < len(points); idx++ {
		tmpSlopes := make(map[string]int)
		tmpMax := 0
		double := 0
		for inner := 0; inner < idx; inner++ {
			x1, y1, x2, y2 := points[inner][0], points[inner][1], points[idx][0], points[idx][1]
			if x1 == x2 && y1 == y2 {
				double++
				continue
			}
			k := slope(x1, y1, x2, y2)
			if _, ok := tmpSlopes[k]; !ok {
				tmpSlopes[k] = 1
			}
			tmpSlopes[k]++
			if tmpSlopes[k] > _max {
				_max = tmpSlopes[k]
			}
		}

		if tmpMax == 0 {
			tmpMax = 1
		}
		tmpMax += double
		if tmpMax > _max {
			_max = tmpMax
		}
	}
	return _max
}

func slope(x1, y1, x2, y2 int) string {
	xx := x2 - x1
	yy := y2 - y1
	if xx == 0 {
		return fmt.Sprintf("%d", x1)
	}
	if yy == 0 {
		return fmt.Sprintf("%d", y1)
	}
	flag := xx*yy < 0
	if xx < 0 {
		xx = -xx
	}
	if yy < 0 {
		yy = -yy
	}
	k := gcd(xx, yy)
	yy /= k
	xx /= k
	s := fmt.Sprintf("%d/%d", yy, xx)
	if flag {
		return "-" + s
	}
	return s
}

// 斜率问题，考虑午饭整除的情况
func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
