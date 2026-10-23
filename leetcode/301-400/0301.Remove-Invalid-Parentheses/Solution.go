package Solution

import (
	"maps"
	"slices"
)

func Solution(s string) []string {
	//感觉需要先判断最少的移除次数
	left, right := 0, 0
	for i := range s {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			if left == 0 {
				right++
			} else {
				left--
			}
		}
	}

	var dfs func(int, int, int, int, string)
	retMap := map[string]struct{}{}

	dfs = func(index, lRm, rRm, balance int, path string) {
		if balance < 0 {
			return
		}

		if index == len(s) {
			if lRm == 0 && rRm == 0 && balance == 0 {
				retMap[path] = struct{}{}
				return
			}
			return
		}

		char := s[index]
		if char == '(' && lRm > 0 {
			dfs(index+1, lRm-1, rRm, balance, path)
		}
		if char == ')' && rRm > 0 {
			dfs(index+1, lRm, rRm-1, balance, path)
		}
		if char == '(' {
			dfs(index+1, lRm, rRm, balance+1, path+string(char))
		} else if char == ')' {
			dfs(index+1, lRm, rRm, balance-1, path+string(char))
		} else {
			dfs(index+1, lRm, rRm, balance, path+string(char))
		}
	}
	dfs(0, left, right, 0, "")
	return slices.Collect(maps.Keys(retMap))
}
