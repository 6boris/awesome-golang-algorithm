package Solution

func Solution(nums []int, k int) bool {
	hash := make(map[int]int)
	for i, num := range nums {
		if v, ok := hash[num]; ok {
			if i-v <= k {
				return true
			}
		}
		hash[num] = i
	}
	return false
}
