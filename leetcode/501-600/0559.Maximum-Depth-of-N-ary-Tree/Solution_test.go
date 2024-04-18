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
		inputs *Node
		expect int
	}{
		{"TestCase1", &Node{
			Val: 1,
			Children: []*Node{
				{Val: 3, Children: []*Node{{Val: 5}, {Val: 6}}}, {Val: 2}, {Val: 4},
			},
		}, 3},
		{"TestCase2", &Node{
			Val: 1,
			Children: []*Node{
				{Val: 2},
				{Val: 3, Children: []*Node{{Val: 6}, {
					Val: 7, Children: []*Node{{Val: 11, Children: []*Node{{Val: 14}}}},
				}}},
				{Val: 4, Children: []*Node{{Val: 8, Children: []*Node{{Val: 12}}}}},
				{Val: 5, Children: []*Node{{Val: 9, Children: []*Node{{Val: 13}}}, {Val: 10}}},
			},
		}, 5},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
