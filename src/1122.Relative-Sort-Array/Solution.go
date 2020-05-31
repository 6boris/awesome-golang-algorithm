package Solution

import "sort"

func Solution(arr1 []int, arr2 []int) []int {
	hash := make(map[int]int)
	for _, val := range arr1 {
		hash[val]++
	}
	res, others := make([]int, 0), make([]int, 0)
	for _, val := range arr2 {
		if n, ok := hash[val]; ok {
			for i := 0; i < n; i++ {
				res = append(res, val)
			}
		}
		delete(hash, val)
	}
	for k, v := range hash {
		for i := 0; i < v; i++ {
			others = append(others, k)
		}
	}
	sort.Ints(others)
	res = append(res, others...)
	return res
}
