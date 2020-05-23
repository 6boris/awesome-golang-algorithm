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
