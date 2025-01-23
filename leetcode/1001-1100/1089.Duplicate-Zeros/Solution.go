package Solution

func Solution(arr []int) {
	n := len(arr)
	if n > 1 {
		for i := 0; i < n; i++ {
			if arr[i] == 0 {
				for j := n - 1; j > i; j-- {
					arr[j] = arr[j-1]
				}
				i++
			}
		}
	}
}

func Solution1(arr []int) {
	tmp := make([]int, len(arr))
	index, i := 0, 0
	for ; i < len(arr) && index < len(arr); i++ {
		if arr[i] != 0 {
			tmp[index] = arr[i]
			index++
			continue
		}
		index += 2
	}
	copy(arr, tmp)
}
