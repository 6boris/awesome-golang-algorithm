package Solution

import "strings"

type pair331 struct {
	sharp, index int
}

func Solution(preorder string) bool {
	if preorder == "#" {
		return true
	}
	items := strings.Split(preorder, ",")
	l := len(items)
	if l < 3 {
		return false
	}
	pairs := make([]pair331, 0)
	for i := 0; i < l; i++ {
		if items[i] != "#" {
			pairs = append(pairs, pair331{index: i})
			continue
		}
		last := len(pairs) - 1
		if last == -1 {
			return false
		}
		if items[pairs[last].index] == "#" || pairs[last].sharp >= 2 {
			pairs = append(pairs, pair331{index: i})
			continue
		}
		pairs[last].sharp++
	}

	stack := []int{-1, -1}
	for _, pair := range pairs {
		if pair.sharp == 1 {
			stack = append(stack, pair.index)
			continue
		}
		if pair.sharp == 0 && items[pair.index] != "#" {
			stack = append(stack, pair.index, pair.index)
			continue
		}
		ll := len(stack)
		if ll == 0 {
			break
		}
		for ll >= 2 && stack[ll-1] != stack[ll-2] {
			ll--
		}
		stack = stack[:ll-1]
	}
	return len(stack) == 1
}
