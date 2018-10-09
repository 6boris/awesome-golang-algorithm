package Solution

import "fmt"

//	解法1：水平暴力扫描
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := 999999
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])

		}
	}

	tmp := strs[0]
	for i := 1; i < minLen; i++ {
		j := 0
		for j <= minLen && tmp[j] == strs[i][j] {
			j++
		}
		fmt.Println(tmp)
		fmt.Println(j)
		tmp = tmp[:j]
	}

	return tmp
}
