package Solution

func Solution(arr []int) bool {
	for i := 0; i < len(arr)-2; i++ {
		if arr[i]&1 == 1 && arr[i+1]&1 == 1 && arr[i+2]&1 == 1 {
			return true
		}
	}
	return false
}
