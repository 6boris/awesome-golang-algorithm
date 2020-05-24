package Solution

func Solution(nums []int) []int {
	res := make([]int, 0)
	for _, val1 := range nums {
		c := 0
		for _, val2 := range nums {
			if val2 < val1 {
				c++
			}
		}
		res = append(res, c)
	}
	return res
}
