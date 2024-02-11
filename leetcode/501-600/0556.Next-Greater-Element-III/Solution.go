package Solution

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func Solution(n int) int {
	bs := []byte(fmt.Sprint(n))
	l := len(bs)
	index := l - 1
	for ; index > 0; index-- {
		if bs[index] > bs[index-1] {
			break
		}
	}
	if index == 0 {
		return -1
	}

	idx := l - 1
	for ; idx >= index; idx-- {
		if bs[index-1] < bs[idx] {
			break
		}
	}
	bs[index-1], bs[idx] = bs[idx], bs[index-1]
	part := bs[index:]
	sort.Slice(part, func(i, j int) bool {
		return part[i] < part[j]
	})

	n1, _ := strconv.Atoi(string(bs))
	if n1 > math.MaxInt32 {
		return -1
	}
	return n1
}
