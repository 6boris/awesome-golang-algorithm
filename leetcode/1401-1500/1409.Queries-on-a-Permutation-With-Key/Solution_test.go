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
		queries []int
		m       int
		expect  []int
	}{
		{"TestCase", []int{3, 1, 2, 1}, 5, []int{2, 1, 2, 1}},
		{"TestCase", []int{4, 1, 2, 2}, 4, []int{3, 1, 2, 0}},
		{"TestCase", []int{7, 5, 5, 8, 3}, 8, []int{6, 5, 0, 7, 5}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.queries, c.m)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.queries, c.m)
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
