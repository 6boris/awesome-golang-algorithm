package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例

	nodes := [6]*Node1{
		{Val: 1},
		{Val: 2},
		{Val: 3},
		{Val: 4},
		{Val: 5},
		{Val: 7},
	}
	nodes[0].Left = nodes[1]
	nodes[0].Right = nodes[2]
	nodes[1].Left = nodes[3]
	nodes[1].Right = nodes[4]
	nodes[2].Right = nodes[5]
	nodes[1].Next = nodes[2]
	nodes[3].Next = nodes[4]
	nodes[4].Next = nodes[5]
	cases := []struct {
		name   string
		inputs *Node1
		expect *Node1
	}{
		{"TestCase1", &Node1{
			Val: 1,
			Left: &Node1{
				Val:   2,
				Left:  &Node1{Val: 4},
				Right: &Node1{Val: 5},
			},
			Right: &Node1{
				Val:   3,
				Right: &Node1{Val: 7},
			},
		}, nodes[0]},
		{"TestCase2", nil, nil},
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
