package Solution

func Solution(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	queue := make([]int, 0)
	toBeUsed := make(map[int]struct{})
	for _, b := range initialBoxes {
		if status[b] == 1 {
			queue = append(queue, b)
			continue
		}
		toBeUsed[b] = struct{}{}
	}

	ans := 0
	for len(queue) > 0 {
		nq := make([]int, 0)
		for _, b := range queue {
			ans += candies[b]
			for _, unlock := range keys[b] {
				status[unlock] = 1
				if _, ok := toBeUsed[unlock]; ok {
					nq = append(nq, unlock)
					delete(toBeUsed, unlock)
				}
			}

			for _, box := range containedBoxes[b] {
				if status[box] == 1 {
					nq = append(nq, box)
					continue
				}
				toBeUsed[box] = struct{}{}
			}

		}
		queue = nq
	}
	return ans
}
