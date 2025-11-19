package Solution

func Solution(skill []int, mana []int) int64 {
	N, M := len(skill), len(mana)
	prevPotion := make([]int64, N)
	for potion := 0; potion < M; potion++ {
		var startTime int64
		if potion > 0 {
			// Need to find the min time in between these two to start
			start, end := prevPotion[0], prevPotion[N-1]
			for start <= end {
				mid := start + (end-start)/2
				if canCompletePotion(mid, potion, prevPotion, skill, mana) {
					startTime = mid
					end = mid - 1
				} else {
					start = mid + 1
				}
			}
		}
		for wiz := 0; wiz < N; wiz++ {
			prevPotion[wiz] = startTime + int64(skill[wiz])*int64(mana[potion])
			startTime = prevPotion[wiz]
		}
	}
	return prevPotion[N-1]
}

func canCompletePotion(startTime int64, potionCount int, prevPotion []int64, skill, mana []int) bool {
	for wizard := 0; wizard < len(skill); wizard++ {
		doneTime := startTime + int64(skill[wizard])*int64(mana[potionCount])
		// If the done time for the previous potion is greater than this one
		// it means we cannot hand off this potion to the next wizard in a sequential
		// manner.
		if prevPotion[wizard] > startTime {
			return false
		}
		startTime = doneTime
	}
	// As long as all potions are done after the previous, we are ok.
	return true
}
