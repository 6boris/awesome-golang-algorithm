package Solution

//	二分查找
func minNumberInRotateArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	first, last := 0, len(arr)-1

	for first < last {
		if arr[first] < arr[last] {
			return arr[first]
		}

		mid := first + ((last - first) >> 1)
		if arr[mid] > arr[last] {
			first = mid + 1
		} else if arr[mid] < arr[last] {
			last = mid
		} else {
			last -= 1
		}
	}

	return arr[first]
}
