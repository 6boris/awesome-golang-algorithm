package Solution

import (
	"strings"
)

type trieNode2227 struct {
	child [26]*trieNode2227
	end   bool
}

func buildTrieNode2227(dic []string) *trieNode2227 {
	root := &trieNode2227{
		child: [26]*trieNode2227{},
	}
	for _, word := range dic {
		walker := root
		for _, b := range []byte(word) {
			index := b - 'a'
			if walker.child[index] == nil {
				walker.child[index] = &trieNode2227{}
			}
			walker = walker.child[index]
		}
		walker.end = true
	}
	return root
}

type Encrypter struct {
	tree         *trieNode2227
	keys         []byte
	keysMap      map[byte]int
	values       []string
	valuesMapper map[string][]int
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	tree := buildTrieNode2227(dictionary)
	e := Encrypter{
		tree:         tree,
		keys:         keys,
		keysMap:      make(map[byte]int),
		values:       values,
		valuesMapper: map[string][]int{},
	}
	for i, b := range keys {
		e.keysMap[b] = i
	}
	for i, str := range values {
		if _, ok := e.valuesMapper[str]; !ok {
			e.valuesMapper[str] = []int{}
		}
		e.valuesMapper[str] = append(e.valuesMapper[str], i)
	}

	return e
}

func (this *Encrypter) Encrypt(word1 string) string {
	buf := strings.Builder{}
	for _, b := range []byte(word1) {
		index, ok := this.keysMap[b]
		if !ok {
			return ""
		}

		buf.WriteString(this.values[index])
	}
	return buf.String()
}

func (this *Encrypter) search(word2 string, index int, tree *trieNode2227) int {
	if index >= len(word2) {
		if tree.end {
			return 1
		}
		return 0
	}
	// ab -> c, c
	cur := word2[index : index+2]
	var ret int
	for _, i := range this.valuesMapper[cur] {
		childIndex := this.keys[i] - 'a'
		if tree.child[childIndex] == nil {
			continue
		}
		ret += this.search(word2, index+2, tree.child[childIndex])
	}
	return ret
}
func (this *Encrypter) Decrypt(word2 string) int {
	return this.search(word2, 0, this.tree)
}

func Solution(keys []byte, values []string, dictionary []string, encryptStr, decryptStr string) []any {
	c := Constructor(keys, values, dictionary)
	var ret []any
	ret = append(ret, c.Encrypt(encryptStr))
	ret = append(ret, c.Decrypt(decryptStr))
	return ret
}
