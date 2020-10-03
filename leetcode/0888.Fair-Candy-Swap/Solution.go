package Solution

func Solution(A []int, B []int) []int {
	d := (sum(A) - sum(B)) / 2
	hash := make(map[int]struct{}, len(A))
	for _, val := range A {
		hash[val] = struct{}{}
	}
	for _, val := range B {
		if _, ok := hash[val+d]; ok {
			return []int{val + d, val}
		}
	}
	return []int{}
}

func sum(nums []int) int {
	var ret int
	for _, num := range nums {
		ret += num
	}
	return ret
}
