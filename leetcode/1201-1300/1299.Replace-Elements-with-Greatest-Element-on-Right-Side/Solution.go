package Solution

func Solution(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		t := arr[i]
		arr[i] = max
		if t > max {
			max = t
		}
	}
	return arr
}
