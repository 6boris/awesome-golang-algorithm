package Solution

import "sort"

func Solution(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for i := range asteroids {
		if mass < asteroids[i] {
			return false
		}
		mass += asteroids[i]
	}

	return true
}
