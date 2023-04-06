package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		s1, s2 string
		expect bool
	}{
		{"TestCacse 1", "great", "rgeat", true},
		{"TestCacse 2", "abcde", "cadbd", false},
		{"TestCacse 3", "a", "a", true},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Solution(c.s1, c.s2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, ret, c.s1, c.s2)
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
