package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		allowed string
		words   []string
		expect  int
	}{
		{"TestCase1", "ab", []string{"ad", "bd", "aaab", "baa", "badab"}, 2},
		{"TestCase2", "abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}, 7},
		{"TestCase3", "cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}, 4},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.allowed, c.words)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.allowed, c.words)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
