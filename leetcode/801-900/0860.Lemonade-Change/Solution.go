package Solution

func Solution(bills []int) bool {
	change := [5]int{}
	for _, cost := range bills {
		idx := cost / 5

		if cost == 5 {
			change[idx]++
			continue
		}

		diff := cost - 5
		need10 := diff / 10
		if change[2] > 0 {
			if change[2] >= need10 {
				change[2] -= need10
				diff -= need10 * 10
			} else {
				diff -= change[2] * 10
				change[2] = 0
			}
		}
		need5 := diff / 5
		if need5 != 0 && change[1] < need5 {
			return false
		}
		change[1] -= need5
		change[idx]++
	}
	return true
}
