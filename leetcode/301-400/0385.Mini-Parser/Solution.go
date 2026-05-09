package Solution

import (
	"fmt"
	"strconv"
	"strings"
)

type NestedInteger struct {
	Value int
	IsNum bool
	Child []NestedInteger
}

func (n NestedInteger) IsInteger() bool {
	return n.IsNum
}

func (n NestedInteger) GetInteger() int {
	return n.Value
}

func (n *NestedInteger) SetInteger(value int) {
	n.Value = value
	n.IsNum = true
}

func (n *NestedInteger) Add(elem NestedInteger) {
	n.Child = append(n.Child, elem)
}

func (n NestedInteger) GetList() []*NestedInteger {
	ret := make([]*NestedInteger, len(n.Child))
	for i := range n.Child {
		ret[i] = &n.Child[i]
	}
	return ret
}

func (n NestedInteger) serialize() string {
	if n.IsNum {
		return fmt.Sprintf("%d", n.Value)
	}
	buf := strings.Builder{}
	buf.WriteByte('[')
	size := len(n.Child)
	if size > 0 {
		for i := 0; i < size-1; i++ {
			buf.WriteString(n.Child[i].serialize() + ",")
		}
		buf.WriteString(n.Child[size-1].serialize())
	}
	buf.WriteByte(']')

	return buf.String()
}

func Solution(s string) *NestedInteger {
	if len(s) == 0 {
		return nil
	}

	root := &NestedInteger{}
	if s[0] != '[' {
		num, _ := strconv.Atoi(s)
		root.SetInteger(num)
		return root
	}

	numStart := -1
	stack := []*NestedInteger{root}
	// "[123,[456,[789]]]"
	for i := 1; i < len(s); i++ {
		if s[i] == '-' || (s[i] >= '0' && s[i] <= '9') {
			if numStart == -1 {
				numStart = i
			}
			continue
		}
		if s[i] == ',' {
			if numStart != -1 {
				num, _ := strconv.Atoi(s[numStart:i])
				node := NestedInteger{}
				node.SetInteger(num)
				stack[len(stack)-1].Add(node)
				numStart = -1
			}
			continue
		}
		// root/ node1 /node2
		if s[i] == '[' {
			node := &NestedInteger{}
			stack = append(stack, node)
			continue
		}
		if s[i] == ']' {
			if numStart != -1 {
				num, _ := strconv.Atoi(s[numStart:i])
				node := NestedInteger{}
				node.SetInteger(num)
				stack[len(stack)-1].Add(node)
				numStart = -1
			}
			l := len(stack)
			top := stack[l-1]
			stack = stack[:l-1]
			if len(stack) > 0 {
				stack[len(stack)-1].Add(*top)
			}
		}
	}
	return root
}
