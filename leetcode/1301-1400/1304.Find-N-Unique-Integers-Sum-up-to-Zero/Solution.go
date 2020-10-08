package Solution

func Solution(n int) []int {
	res := make([]int, 0)
	if n%2 == 0 {
		for i := 0; i < n/2; i++ {
			res = append(res, -i-1)
			res = append(res, i+1)
		}
	} else {
		for i := 0; i < n/2; i++ {
			res = append(res, -i-1)
			res = append(res, i+1)
		}
		res = append(res, 0)
	}
	return res
}
