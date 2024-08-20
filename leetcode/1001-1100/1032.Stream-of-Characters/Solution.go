package Solution

type trie1032 struct {
	end   bool
	child [26]*trie1032
}

type StreamChecker struct {
	// trie
	i  int
	ml int
	q  []byte // max length
	t  *trie1032
}

func Constructor1032(words []string) StreamChecker {
	s := StreamChecker{
		t: &trie1032{},
	}
	ml := 0
	for _, w := range words {
		ml = max(ml, len(w))
		s.fill(w)
	}

	s.q = make([]byte, ml)
	s.ml = ml
	return s
}

func (this *StreamChecker) fill(word string) {
	root := this.t
	for i := len(word) - 1; i >= 0; i-- {
		b := word[i]
		if root.child[b-'a'] == nil {
			root.child[b-'a'] = &trie1032{}
		}
		if i == 0 {
			root.child[b-'a'].end = true
		}
		root = root.child[b-'a']
	}
}

func (this *StreamChecker) Query(letter byte) bool {
	if this.i == this.ml {
		next := make([]byte, this.ml)
		copy(next, this.q[1:])
		this.i = this.ml - 1
		this.q = next
	}
	this.q[this.i] = letter
	this.i++
	root := this.t
	for i := this.i - 1; i >= 0; i-- {
		cur := this.q[i] - 'a'
		if root.child[cur] == nil {
			return false
		}
		if root.child[cur].end {
			return true
		}
		root = root.child[cur]
	}
	return false
}

func Solution(words []string, query []byte) []bool {
	ans := make([]bool, len(query))
	c := Constructor1032(words)
	for i := range query {
		ans[i] = c.Query(query[i])
	}
	return ans
}
