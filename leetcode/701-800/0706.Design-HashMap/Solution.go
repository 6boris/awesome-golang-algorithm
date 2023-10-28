package Solution

type MyHashMap struct {
	// e
	data map[int]int
}

func Constructor706() MyHashMap {
	return MyHashMap{data: map[int]int{}}
}

func (this *MyHashMap) Put(key int, value int) {
	this.data[key] = value
}

func (this *MyHashMap) Get(key int) int {
	if v, ok := this.data[key]; ok {
		return v
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this.data, key)
}

type operation struct {
	name string
	k, v int
}

func Solution(opts []operation) []int {
	c := Constructor706()
	ans := make([]int, 0)
	for _, opt := range opts {
		if opt.name == "put" {
			c.Put(opt.k, opt.v)
			continue
		}
		if opt.name == "remove" {
			c.Remove(opt.k)
			continue
		}
		ans = append(ans, c.Get(opt.k))
	}
	return ans
}
