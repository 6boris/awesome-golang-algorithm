package Solution

import (
	"sort"
	"strings"
)

type filePathNode struct {
	end  bool
	data map[string]*filePathNode
}

func insertFolder(root *filePathNode, path string) bool {
	walker := root
	paths := strings.Split(path, "/")
	for i := 1; i < len(paths); i++ {
		p := paths[i]
		_, ok := walker.data[p]
		if !ok {
			walker.data[p] = &filePathNode{end: false, data: make(map[string]*filePathNode)}
		}
		if i == len(paths)-1 {
			walker.data[p].end = true
			continue
		}
		if walker.data[p].end {
			return false
		}
		walker = walker.data[p]
	}
	return true
}

func Solution(folder []string) []string {
	sort.Slice(folder, func(i, j int) bool {
		return len(folder[i]) < len(folder[j])
	})

	root := &filePathNode{end: false, data: make(map[string]*filePathNode)}
	ans := make([]string, 0)
	for _, str := range folder {
		if insertFolder(root, str) {
			ans = append(ans, str)
		}
	}
	return ans
}
