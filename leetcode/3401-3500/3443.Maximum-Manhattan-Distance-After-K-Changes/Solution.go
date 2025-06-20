package Solution

func Solution(s string, k int) int {
	n1, s1, e1, w1 := 0, 0, 0, 0
	var ans int
	for _, b := range s {
		if b == 'N' {
			n1++
		}
		if b == 'S' {
			s1++
		}
		if b == 'W' {
			w1++
		}
		if b == 'E' {
			e1++
		}
		// 我们只需要将少的移动的多的方向，另外就是，南北，东西变化可以增加，但是南东这种变化没用
		minY := min(n1, s1)
		maxY := max(n1, s1)
		a := min(minY, k)

		minX := min(e1, w1)
		maxX := max(e1, w1)
		b := min(minX, k-a)
		// n1 - (s1-a) + a
		ans = max(ans, maxY-minY+a*2+maxX-minX+b*2)

	}
	return ans
}
