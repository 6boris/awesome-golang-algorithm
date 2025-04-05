package Solution

func merge(nums []int, left, right int) int {
	mid := (right-left)/2 + left
	i, j := left, mid+1
	temp := make([]int, right-left+1)
	k := 0
	reverseOrderPairs := 0
	for ; i <= mid && j <= right; k++ {
		if nums[i] > nums[j] {
			reverseOrderPairs += mid - i + 1
			temp[k] = nums[j]
			j++
			continue
		}
		temp[k] = nums[i]
		i++
	}
	for ; i <= mid; i, k = i+1, k+1 {
		temp[k] = nums[i]

	}
	for ; j <= right; j, k = j+1, k+1 {
		temp[k] = nums[j]
	}
	for ; k > 0; k-- {
		nums[left+k-1] = temp[k-1]
	}
	return reverseOrderPairs
}
func mergeSort(nums []int, left, right int) int {
	if left < right {
		mid := (right-left)/2 + left
		return mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right) + merge(nums, left, right)
	}
	return 0
}

func Solution(nums []int) bool {
	local := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			local++
		}
	}
	global := mergeSort(nums, 0, len(nums)-1)
	return local == global
}
