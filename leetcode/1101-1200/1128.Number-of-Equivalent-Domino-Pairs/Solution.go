package Solution

func Solution(dominoes [][]int) int {
	n := len(dominoes)
	if n == 1 {
		return 0
	}
	hash, pairs := make(map[int]int), 0
	for _, val := range dominoes {
		newVal := changeVal(val)
		pairs += hash[newVal]
		hash[newVal]++
	}
	return pairs
}

func changeVal(val []int) int {
	if val[0] > val[1] {
		val[0], val[1] = val[1], val[0]
	}
	return val[0]*10 + val[1]
}
