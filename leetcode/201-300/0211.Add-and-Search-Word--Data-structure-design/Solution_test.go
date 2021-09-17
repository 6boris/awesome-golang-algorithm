package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	wordDictionary.Search("pad") // return False
	wordDictionary.Search("bad") // return True
	wordDictionary.Search(".ad") // return True
	wordDictionary.Search("b..") // return True
}
