package Solution

func Solution(nums []int, k int) {
	n := len(nums)
	if k == 0 || k%n == 0 {
		return
	}
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(arr []int, x, y int) {
	for i, j := x, y; i <= j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
