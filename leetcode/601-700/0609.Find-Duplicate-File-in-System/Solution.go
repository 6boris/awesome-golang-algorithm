package Solution

import (
	"path/filepath"
	"strings"
)

func Solution(paths []string) [][]string {
	cache := make(map[string][]string)
	for _, path := range paths {
		files := strings.Split(path, " ")
		baseDir := files[0]
		for _, file := range files[1:] {
			nameAndContent := strings.Split(file, ".txt")
			name, content := nameAndContent[0], nameAndContent[1]
			if _, ok := cache[content]; !ok {
				cache[content] = make([]string, 0)
			}
			cache[content] = append(cache[content], filepath.Join(baseDir, name+".txt"))
		}
	}
	ans := make([][]string, 0)
	for _, f := range cache {
		if len(f) >= 2 {
			ans = append(ans, f)
		}
	}
	return ans
}
