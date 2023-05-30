package Solution

type op struct {
	name string
	v    int
}
type MyHashSet struct {
	//m [1000][1001]uint8
	m [1000001]byte
}

func Constructor705() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	this.m[key] = 1
}

func (this *MyHashSet) Remove(key int) {
	this.m[key] = 0
}

func (this *MyHashSet) Contains(key int) bool {
	return this.m[key] == 1
}

func Solution(ops []op) []bool {
	c := Constructor705()
	ans := make([]bool, 0)
	for _, o := range ops {
		if o.name == "add" {
			c.Add(o.v)
			continue
		}
		if o.name == "del" {
			c.Remove(o.v)
			continue
		}

		ans = append(ans, c.Contains(o.v))
	}
	return ans
}
