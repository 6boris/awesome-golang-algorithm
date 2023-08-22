package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name          string
		numCourses    int
		prerequisties [][]int
		expect        []int
	}{
		{"TestCase1", 2, [][]int{{1, 0}}, []int{0, 1}},
		{"TestCase2", 4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 1, 2, 3}},
		{"TestCase3", 1, nil, []int{0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.numCourses, c.prerequisties)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.numCourses, c.prerequisties)
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
