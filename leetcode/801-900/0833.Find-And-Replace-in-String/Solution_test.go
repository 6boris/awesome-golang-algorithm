package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name             string
		s                string
		indices          []int
		sources, targets []string
		expect           string
	}{
		{"TestCase1", "abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}, "eeebffff"},
		{"TestCase2", "abcd", []int{0, 2}, []string{"ab", "ed"}, []string{"eee", "ffff"}, "eeecd"},
		{"TestCase3", "abcd", []int{0, 3, 2, 1}, []string{"a", "d", "c", "b"}, []string{"d", "a", "b", "c"}, "dcba"},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.s, c.indices, c.sources, c.targets)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.s, c.indices, c.sources, c.targets)
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
