package Solution

import (
	"errors"
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Stack struct {
	lock  sync.Mutex
	nodes []*TreeNode
}

func NewStack() *Stack {
	return &Stack{sync.Mutex{}, []*TreeNode{}}
}

func (s *Stack) Push(node *TreeNode) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.nodes = append(s.nodes, node)
}

func (s *Stack) Pop() (*TreeNode, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.nodes)
	if l == 0 {
		return nil, errors.New("Empty")
	}

	node := s.nodes[l-1]
	s.nodes = s.nodes[:l-1]

	return node, nil
}

func postorderTraversal(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	stack1 := NewStack()
	stack2 := NewStack()

	stack1.Push(root)
	for {
		node, err := stack1.Pop()
		if err != nil {
			break
		}

		stack2.Push(node)
		if node.Left != nil {
			stack1.Push(node.Left)
		}

		if node.Right != nil {
			stack1.Push(node.Right)
		}
	}

	for {
		node, err := stack2.Pop()
		if err != nil {
			break
		}

		result = append(result, node.Val)
	}

	return result
}
