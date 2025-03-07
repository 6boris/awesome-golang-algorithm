package Solution

import "strings"

func Solution(target string) string {
	pos := [26][2]int{
		{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
		{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4},
		{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4},
		{3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4},
		{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4},
		{5, 0},
	}
	buf := strings.Builder{}
	x, y := 0, 0
	for _, b := range target {
		bp := pos[b-'a']
		if bp[0] == x && bp[1] == y {
			buf.WriteByte('!')
			continue
		}
		dx, dy := bp[0]-x, bp[1]-y
		sx, sy := x, y
		x, y = bp[0], bp[1]
		if sx == 5 && sy == 0 {
			if dx < 0 {
				for ; dx < 0; dx++ {
					buf.WriteByte('U')
				}
			} else {
				for ; dx > 0; dx-- {
					buf.WriteByte('D')
				}
			}
			if dy < 0 {
				for ; dy < 0; dy++ {
					buf.WriteByte('L')
				}
			} else {
				for ; dy > 0; dy-- {
					buf.WriteByte('R')
				}
			}
			buf.WriteByte('!')
			continue
		}
		if dy < 0 {
			for ; dy < 0; dy++ {
				buf.WriteByte('L')
			}
		} else {
			for ; dy > 0; dy-- {
				buf.WriteByte('R')
			}
		}
		if dx < 0 {
			for ; dx < 0; dx++ {
				buf.WriteByte('U')
			}
		} else {
			for ; dx > 0; dx-- {
				buf.WriteByte('D')
			}
		}
		buf.WriteByte('!')
	}
	return buf.String()
}
