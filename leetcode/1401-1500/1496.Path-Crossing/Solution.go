package Solution

func Solution(path string) bool {
	x, y := 0, 0
	cache := make(map[[2]int]struct{})
	cache[[2]int{x, y}] = struct{}{}
	for _, b := range path {
		switch b {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}
		if _, ok := cache[[2]int{x, y}]; ok {
			return true
		}
		cache[[2]int{x, y}] = struct{}{}
	}
	return false
}
