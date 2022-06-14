package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name         string
		word1, word2 string
		expect       int
	}{
		{"TestCase1", "sea", "eat", 2},
		{"TestCase2", "leetcode", "etco", 4},
		{"TestCase3", "a", "a", 0},
		{"TestCase4", "aaaaaaa", "aaaa", 3},
		{"TestCase4", "a", "b", 2},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.word1, c.word2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.word1, c.word2)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
