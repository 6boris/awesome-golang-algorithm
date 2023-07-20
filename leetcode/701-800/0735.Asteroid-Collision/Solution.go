package Solution

func Solution(asteroids []int) []int {
	latestIndex := 0

	for idx := 1; idx < len(asteroids); idx++ {
		if latestIndex < 0 || asteroids[latestIndex] < 0 || asteroids[idx] > 0 {
			latestIndex++
			asteroids[latestIndex] = asteroids[idx]
			continue
		}
		// latest > 0 && asteroids[idx] < 0
		add := true
		rev := -asteroids[idx]
		for latestIndex >= 0 {
			if asteroids[latestIndex] < 0 {
				break
			}

			if asteroids[latestIndex] == rev {
				latestIndex--
				add = false
				break
			}
			if asteroids[latestIndex] > rev {
				add = false
				break
			}
			latestIndex--
		}

		if add {
			latestIndex++
			asteroids[latestIndex] = asteroids[idx]
		}
	}
	return asteroids[:latestIndex+1]
}
