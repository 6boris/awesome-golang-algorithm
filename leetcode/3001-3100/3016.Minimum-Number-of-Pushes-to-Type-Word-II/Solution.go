package Solution

import "sort"

type char struct {
	ch byte
	c  int
}

func Solution(word string) int {
	chars := [26]int{}
	for _, b := range word {
		chars[b-'a']++
	}
	list := make([]char, 0)
	for i := 0; i < 26; i++ {
		if chars[i] != 0 {
			list = append(list, char{ch: uint8(i) + 'a', c: chars[i]})
		}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].c > list[j].c
	})
	// 2-9
	cur := 2
	loop := 1
	need := 0
	for _, item := range list {
		if cur == 10 {
			cur = 2
			loop++
		}
		need += item.c * loop
		cur++
	}
	return need
}
