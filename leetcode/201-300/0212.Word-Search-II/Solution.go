package Solution

type trieNode212 struct {
	child [26]*trieNode212
	word  string
}

func buildTrieNode212(words []string) *trieNode212 {
	root := &trieNode212{}
	for _, word := range words {
		walker := root
		// a, b
		for i, b := range word {
			index := b - 'a'
			if walker.child[index] == nil {
				walker.child[index] = &trieNode212{}
			}
			walker = walker.child[index]
			if i == len(word)-1 {
				walker.word = word
			}
		}
	}
	return root
}
func Solution(board [][]byte, words []string) []string {
	rows, cols := len(board), len(board[0])
	tree := buildTrieNode212(words)
	var (
		search func(int, int, *trieNode212)
		dirs   = [][]int{
			{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		}
	)
	found := map[string]struct{}{}
	search = func(x, y int, tree *trieNode212) {
		if x < 0 || x >= rows || y < 0 || y >= cols || board[x][y] == '|' {
			return
		}
		index := board[x][y] - 'a'
		node := tree.child[index]
		if node == nil {
			return
		}
		if node.word != "" {
			found[node.word] = struct{}{}
		}
		source := board[x][y]
		board[x][y] = '|'
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			walker := node
			search(nx, ny, walker)
		}
		board[x][y] = source
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			search(r, c, tree)
		}
	}
	var ret []string
	for key := range found {
		ret = append(ret, key)
	}
	return ret
}
