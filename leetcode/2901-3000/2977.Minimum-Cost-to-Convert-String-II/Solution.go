package Solution

import "math"

type Trie struct {
	child [26]*Trie
	id    int
}

func newTrie() *Trie {
	return &Trie{id: -1}
}

func add(node *Trie, word string, index *int) int {
	for _, ch := range word {
		i := ch - 'a'
		if node.child[i] == nil {
			node.child[i] = newTrie()
		}
		node = node.child[i]
	}
	if node.id == -1 {
		*index++
		node.id = *index
	}
	return node.id
}

func update(x *int64, y int64) {
	if *x == -1 || y < *x {
		*x = y
	}
}

func Solution(source string, target string, original []string, changed []string, cost []int) int64 {
	n := len(source)
	m := len(original)
	root := newTrie()

	p := -1
	nodeCount := m * 2
	G := make([][]int, nodeCount)
	for i := range G {
		G[i] = make([]int, nodeCount)
		for j := range G[i] {
			G[i][j] = math.MaxInt32 / 2
		}
		G[i][i] = 0
	}

	for i := 0; i < m; i++ {
		x := add(root, original[i], &p)
		y := add(root, changed[i], &p)
		G[x][y] = min(G[x][y], cost[i])
	}

	size := p + 1
	for k := 0; k < size; k++ {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				G[i][j] = min(G[i][j], G[i][k]+G[k][j])
			}
		}
	}

	f := make([]int64, n)
	for i := range f {
		f[i] = -1
	}

	for j := 0; j < n; j++ {
		if j > 0 && f[j-1] == -1 {
			continue
		}

		var base int64
		if j == 0 {
			base = 0
		} else {
			base = f[j-1]
		}

		if source[j] == target[j] {
			update(&f[j], base)
		}

		u, v := root, root
		for i := j; i < n; i++ {
			u = u.child[source[i]-'a']
			v = v.child[target[i]-'a']
			if u == nil || v == nil {
				break
			}
			if u.id != -1 && v.id != -1 && G[u.id][v.id] != math.MaxInt32/2 {
				newVal := base + int64(G[u.id][v.id])
				update(&f[i], newVal)
			}
		}
	}

	return f[n-1]
}
