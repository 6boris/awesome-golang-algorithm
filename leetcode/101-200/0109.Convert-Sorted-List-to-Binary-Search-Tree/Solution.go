package Solution

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solution(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	left, right := head, slow.Next
	mid := slow
	if prev != nil {
		prev.Next = nil
	}
	root := &TreeNode{Val: mid.Val}
	root.Left = Solution(left)
	root.Right = Solution(right)
	return root
}
