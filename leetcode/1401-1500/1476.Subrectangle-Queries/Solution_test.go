package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name      string
		rectangle [][]int
		ops       []op
		expect    []int
	}{
		{"TestCase1", [][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}}, []op{
			{"GetValue", 0, 2, 0, 0, 0},
			{"UpdateSubrectangle", 0, 0, 3, 2, 5},
			{"GetValue", 0, 2, 0, 0, 0},
			{"GetValue", 3, 1, 0, 0, 0},
			{"UpdateSubrectangle", 3, 0, 3, 2, 10},
			{"GetValue", 3, 1, 0, 0, 0},
			{"GetValue", 0, 2, 0, 0, 0},
		}, []int{1, 5, 5, 10, 5}},
		{"TestCase2", [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}, []op{
			{"GetValue", 0, 0, 0, 0, 0},
			{"UpdateSubrectangle", 0, 0, 2, 2, 100},
			{"GetValue", 0, 0, 0, 0, 0},
			{"GetValue", 2, 2, 0, 0, 0},
			{"UpdateSubrectangle", 1, 1, 2, 2, 20},
			{"GetValue", 2, 2, 0, 0, 0},
		}, []int{1, 100, 100, 20}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.rectangle, c.ops)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.rectangle, c.ops)
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
