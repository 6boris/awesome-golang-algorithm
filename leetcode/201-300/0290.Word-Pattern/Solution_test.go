package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name       string
		pattern, s string
		expect     bool
	}{
		{"TestCase1", "abba", "dog cat cat dog", true},
		{"TestCase2", "abba", "dog cat cat fish", false},
		{"TestCase3", "aaaa", "dog cat cat dog", false},
		{"TestCase4", "abba", "dog dog dog dog", false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.pattern, c.s)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.pattern, c.s)
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
