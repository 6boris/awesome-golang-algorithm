package Solution

func Solution(nums []int, k int) int {
	n := len(nums)
	if n < 2 || k < 0 {
		return 0
	}
	hash, c := make(map[int]int), 0
	for _, num := range nums {
		hash[num]++
	}
	if k == 0 {
		for _, v := range hash {
			if v >= 2 {
				c++
			}
		}
	} else {
		for key := range hash {
			if hash[key+k] > 0 {
				c++
			}
		}
	}
	return c
}
