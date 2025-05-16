package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		words  []string
		groups []int
		expect []string
	}{
		{"TestCase1", []string{"bab", "dab", "cab"}, []int{1, 2, 2}, []string{"bab", "dab"}},
		{"TestCase2", []string{"a", "b", "c", "d"}, []int{1, 2, 3, 4}, []string{"a", "b", "c", "d"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.words, c.groups)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.words, c.groups)
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
