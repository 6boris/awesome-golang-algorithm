package Solution

type kingTree struct {
	name  string
	live  bool
	child []*kingTree
}

func (k *kingTree) preOrder(out *[]string) {
	if k == nil {
		return
	}
	if k.live {
		*out = append(*out, k.name)
	}
	for _, c := range k.child {
		c.preOrder(out)
	}
}

// 相当于是n叉树的前序遍历
type ThroneInheritance struct {
	root      *kingTree
	nodeCache map[string]*kingTree
}

func Constructor(kingName string) ThroneInheritance {
	king := &kingTree{
		name:  kingName,
		live:  true,
		child: make([]*kingTree, 0),
	}

	cache := map[string]*kingTree{
		kingName: king,
	}
	ti := ThroneInheritance{
		root:      king,
		nodeCache: cache,
	}
	return ti
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	node := this.nodeCache[parentName]
	childNode := &kingTree{
		name:  childName,
		live:  true,
		child: make([]*kingTree, 0),
	}
	node.child = append(node.child, childNode)
	this.nodeCache[childName] = childNode

}

func (this *ThroneInheritance) Death(name string) {
	this.nodeCache[name].live = false
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	var out []string
	this.root.preOrder(&out)
	return out
}

type operation struct {
	name       string
	parentName string
	childName  string
}

func Solution(kingName string, inputs []operation) [][]string {
	var ret [][]string
	c := Constructor(kingName)
	for _, op := range inputs {
		if op.name == "birth" {
			c.Birth(op.parentName, op.childName)
			continue
		}
		if op.name == "death" {
			c.Death(op.childName)
			continue
		}
		ret = append(ret, c.GetInheritanceOrder())
	}
	return ret
}
