package Solution

func Solution(nums []int) int {
	s := []int{}
	res := 0
	for _, a := range nums {
		for len(s) > 0 && s[len(s)-1] > a {
			s = s[:len(s)-1]
		}
		if a == 0 {
			continue
		}
		if len(s) == 0 || s[len(s)-1] < a {
			res++
			s = append(s, a)
		}
	}
	return res
}
