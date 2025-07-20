package Solution

import (
	"fmt"
	"sort"
	"strings"
)

type Tree struct {
	Value      string
	ValueIndex map[string]int
	Children   []*Tree
	Del        bool
}

func Solution(paths [][]string) [][]string {
	sort.Slice(paths, func(i, j int) bool {
		a, b := len(paths[i]), len(paths[j])
		l := min(a, b)
		for k := range l {
			if paths[i][k] != paths[j][k] {
				return paths[i][k] < paths[j][k]
			}
		}
		return a < b
	})
	// 将所有的子目录序列化成一个字符串。之前有过这样的题，搞成字符串
	root := &Tree{Value: "/", Children: []*Tree{}, ValueIndex: map[string]int{}}
	var buildTree func([]string)
	buildTree = func(p []string) {
		walker := root
		for _, sub := range p {
			index, ok := walker.ValueIndex[sub]
			if ok {
				walker = walker.Children[index]
				continue
			}
			tree := &Tree{
				Value: sub, Children: []*Tree{}, ValueIndex: map[string]int{},
			}
			walker.Children = append(walker.Children, tree)
			walker.ValueIndex[sub] = len(walker.Children) - 1
			walker = tree
		}
	}
	for _, path := range paths {
		buildTree(path)
	}

	var buildNodeFlag func(*Tree) string

	pathToNode := map[string][]*Tree{}
	buildNodeFlag = func(node *Tree) string {
		if node == nil {
			return "#"
		}
		if len(node.Children) == 0 {
			return node.Value + "#"
		}
		sb := strings.Builder{}
		for i, c := range node.Children {
			tmp := buildNodeFlag(c) + fmt.Sprintf("%d", i)
			sb.WriteString(tmp)
		}
		s := sb.String()
		pathToNode[s] = append(pathToNode[s], node)
		return node.Value + s
	}
	buildNodeFlag(root)
	for path, nodes := range pathToNode {
		if len(path) == 0 {
			continue
		}
		if len(nodes) > 1 {
			// 有重复的
			for _, n := range nodes {
				n.Del = true
			}
		}
	}
	var filterPath func([]string) []string
	filterPath = func(cur []string) []string {
		walker := root
		i := 0
		for ; i < len(cur); i++ {
			index := walker.ValueIndex[cur[i]]
			if walker.Children[index].Del {
				break
			}
			walker = walker.Children[index]
		}
		return cur[:i]
	}
	in := map[string]struct{}{}
	var ans [][]string
	for _, path := range paths {
		res := filterPath(path)
		if len(res) > 0 {
			key := strings.Join(res, "/")
			if _, ok := in[key]; ok {
				continue
			}
			ans = append(ans, res)
			in[key] = struct{}{}
		}
	}
	return ans
}
