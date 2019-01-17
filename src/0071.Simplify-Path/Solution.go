package Solution

import (
	"path/filepath"
	"strings"
)

func simplifyPath(path string) string {
	return filepath.Clean(path)
}

func simplifyPath2(path string) string {
	var stack []string
	for _, component := range strings.Split(path, "/") {
		if component == "" || component == "." {
			continue
		}
		if component == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, component)
		}
	}
	return "/" + strings.Join(stack, "/")
}
