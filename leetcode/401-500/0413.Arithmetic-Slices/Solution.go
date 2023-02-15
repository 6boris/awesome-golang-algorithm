package Solution

func Solution(nums []int) int {
	length := len(nums)
	if length < 3 {
		return 0
	}
	diff := make([]int, length-1)
	for i := 0; i < length-1; i++ {
		diff[i] = nums[i+1] - nums[i]
	}
	a := 0
	same := 0
	ans := 0
	for i := 1; i < length-1; i++ {
		if diff[i] == diff[a] {
			same++
			continue
		}
		if same >= 1 {
			ans += (same * (same + 1)) / 2
		}
		same = 0
		a = i
	}
	if same >= 1 {
		ans += (same * (same + 1)) / 2
	}

	return ans
}
