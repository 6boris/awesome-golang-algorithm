package Solution

import (
	"sort"
)

const mod823 = 1000000007

func Solution(arr []int) int {
	ans := 0
	set := make(map[uint64]int)
	for _, n := range arr {
		set[uint64(n)] = 1
	}
	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[i]%arr[j] == 0 {
				tmp := uint64(arr[i]) / uint64(arr[j])
				if count, ok := set[tmp]; ok {
					a := set[uint64(arr[j])]
					set[uint64(arr[i])] += (a * count) % mod823
				}
			}
		}
	}
	for _, c := range set {
		ans = (ans + c) % mod823
	}
	return ans
}
