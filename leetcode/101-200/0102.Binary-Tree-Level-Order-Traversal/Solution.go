package Solution

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func levelOrder_1(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(head *TreeNode, level int) {
		if head == nil {
			return
		}
		if level >= len(ans) {
			ans = append(ans, []int{})
		}
		ans[level] = append(ans[level], head.Val)
		dfs(head.Left, level+1)
		dfs(head.Right, level+1)
	}
	if root != nil {
		ans = append(ans, []int{root.Val})
		dfs(root.Left, 1)
		dfs(root.Right, 1)
	}
	return ans
}

func levelOrder_2(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	// 需要有一个数据结构存储下一层的所有节点
	queue := list.New()
	queue.PushBack(root)
	// 到最后一层时下层节点数为 0
	for queue.Len() > 0 {
		tmp := []int{}
		length := queue.Len()
		// 保存当层节点数量
		for i := 0; i < length; i++ {
			// 检查当前节点到子节点，为遍历下层节点做准备
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			// 记录当前节点值
			tmp = append(tmp, node.Val)
		}
		ans = append(ans, tmp)
	}
	return ans
}

func levelOrder_3(root *TreeNode) [][]int {
	ans, que := make([][]int, 0), make([]*TreeNode, 0)
	if root == nil {
		return ans
	}
	que = append(que, root)
	for len(que) > 0 {
		tmp, length := []int{}, len(que)
		for i := 0; i < length; i++ {
			node := que[0]
			que = que[1:]
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		ans = append(ans, tmp)
	}
	return ans
}
