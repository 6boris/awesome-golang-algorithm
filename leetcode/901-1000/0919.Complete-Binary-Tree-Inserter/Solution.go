package Solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type CBTInserter struct {
	root  *TreeNode
	cache []*TreeNode
}

func Constructor919(root *TreeNode) CBTInserter {
	o := CBTInserter{
		root:  root,
		cache: make([]*TreeNode, 0),
	}
	queue := []*TreeNode{root}
	need := 1
	for len(queue) > 0 {
		next := make([]*TreeNode, 0)
		if len(queue) == need {
			o.cache = queue
		}
		for i := range queue {
			if queue[i].Left != nil {
				next = append(next, queue[i].Left)
			}
			if queue[i].Right != nil {
				next = append(next, queue[i].Right)
			}
		}
		queue = next
		need *= 2
	}
	return o
}

func (this *CBTInserter) Insert(val int) int {
	isLastRight := false
	n := &TreeNode{Val: val}
	ans := 0
	for idx, node := range this.cache {
		if node.Left == nil {
			this.cache[idx].Left = n
			ans = node.Val
			break
		}
		if node.Right == nil {
			this.cache[idx].Right = n
			isLastRight = idx == len(this.cache)-1
			ans = node.Val
			break
		}
	}
	if isLastRight {
		next := make([]*TreeNode, len(this.cache)*2)
		index := 0
		for i := 0; i < len(this.cache); i++ {
			next[index] = this.cache[i].Left
			next[index+1] = this.cache[i].Right
			index += 2
		}
		this.cache = next
	}
	return ans
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

type opt struct {
	name string
	val  int
}

type result struct {
	val  int
	tree *TreeNode
}

func Solution(root *TreeNode, opts []opt) []result {
	ans := make([]result, 0)
	o := Constructor919(root)
	for _, op := range opts {
		if op.name == "insert" {
			item := o.Insert(op.val)
			ans = append(ans, result{val: item})
			continue
		}
		ans = append(ans, result{tree: o.Get_root()})
	}
	return ans
}
