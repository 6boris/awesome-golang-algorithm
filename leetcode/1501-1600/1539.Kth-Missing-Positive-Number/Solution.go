package Solution

func Solution(arr []int, k int) int {
	i := 1
	index := 0
	for ; index < len(arr) && k > 0; i++ {
		if i != arr[index] {
			k--
			continue
		}
		index++
	}
	for ; k > 0; k, i = k-1, i+1 {

	}
	return i - 1
}
