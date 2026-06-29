package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name     string
		patterns []string
		word     string
		expect   int
	}{
		{"TestCase1", []string{"a", "abc", "bc", "d"}, "abc", 3},
		{"TestCase2", []string{"a", "b", "c"}, "aaaaabbbbb", 2},
		{"TestCase3", []string{"a", "a", "a"}, "ab", 3},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.patterns, c.word)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.patterns, c.word)
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
