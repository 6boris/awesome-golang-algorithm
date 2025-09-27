package Solution

import "math"

const epsilon = 1e-9

func canForm(a, b, c []int) (float64, bool) {
	abx := b[0] - a[0]
	aby := b[1] - a[1]
	acx := c[0] - a[0]
	acy := c[1] - a[1]
	area := math.Abs(float64(abx*acy - aby*acx))
	if area > epsilon {
		return area, true
	}
	return 0.0, false
}

func Solution(points [][]int) float64 {
	var ret float64
	l := len(points)
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				area, ok := canForm(points[i], points[j], points[k])
				if ok {
					ret = max(ret, area/2.0)
				}
			}
		}
	}
	return ret
}
