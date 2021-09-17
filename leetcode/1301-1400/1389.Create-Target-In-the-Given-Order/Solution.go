package Solution

func Solution(nums []int, index []int) []int {
	n := len(nums)
	if n == 1 {
		return []int{nums[0]}
	}
	target := make([]int, n)
	for i, val := range index {
		for j := n - 1; j > val; j-- {
			target[j] = target[j-1]
		}
		target[val] = nums[i]
	}
	return target
}
