package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                   string
		numCourses             int
		prerequisties, queries [][]int

		expect []bool
	}{
		{"TestCase1", 2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}, []bool{false, true}},
		{"TestCase2", 2, [][]int{}, [][]int{{0, 1}, {1, 0}}, []bool{false, false}},
		{"TestCase3", 3, [][]int{{1, 2}, {1, 0}, {2, 0}}, [][]int{{1, 0}, {1, 2}}, []bool{true, true}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.numCourses, c.prerequisties, c.queries)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.numCourses, c.prerequisties, c.queries)
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
