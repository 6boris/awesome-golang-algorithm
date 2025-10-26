package Solution

func Solution(total int, cost1 int, cost2 int) int64 {
	var ret int64
	penCount := 0
	for ; penCount*cost1 <= total; penCount++ {
		left := (total-penCount*cost1)/cost2 + 1 // 0
		ret += int64(left)
	}
	return ret
}
