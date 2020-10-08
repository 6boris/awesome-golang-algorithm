package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	trie.Search("apple")   // returns true
	trie.Search("app")     // returns false
	trie.StartsWith("app") // returns true
	trie.Insert("app")
	trie.Search("appl") // returns true
}
