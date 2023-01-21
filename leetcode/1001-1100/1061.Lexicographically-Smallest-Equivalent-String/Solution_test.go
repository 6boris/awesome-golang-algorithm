package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name            string
		s1, s2, baseStr string
		expect          string
	}{
		{"TestCase1", "parker", "morris", "parser", "makkek"},
		{"TestCase2", "hello", "world", "hold", "hdld"},
		{"TestCase3", "leetcode", "programs", "sourcecode", "aauaaaaada"},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.s1, c.s2, c.baseStr)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.s1, c.s2, c.baseStr)
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
