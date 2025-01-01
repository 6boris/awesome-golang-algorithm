package Solution

func Solution(nums []int, pivot int) []int {
	i, j := 0, len(nums)-1
	res := make([]int, len(nums))
	for k := 0; k < len(nums); k++ {
		if nums[k] < pivot {
			res[i] = nums[k]
			i++
		}
	}
	for k := len(nums) - 1; k >= 0; k-- {
		if nums[k] > pivot {
			res[j] = nums[k]
			j--
		}
	}
	for ; i <= j; i++ {
		res[i] = pivot
	}
	return res
}
