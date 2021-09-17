package Solution

func containsDuplicate(nums []int) bool {
	disct := make(map[int]bool, len(nums))

	for _, number := range nums {
		if disct[number] {
			return true
		} else {
			disct[number] = true
		}
	}
	return false
}
