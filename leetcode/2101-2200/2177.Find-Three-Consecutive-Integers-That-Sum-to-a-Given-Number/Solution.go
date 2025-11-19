package Solution

func Solution(num int64) []int64 {
	if num%3 != 0 {
		return []int64{}
	}
	mid := num / 3
	return []int64{mid - 1, mid, mid + 1}
}
