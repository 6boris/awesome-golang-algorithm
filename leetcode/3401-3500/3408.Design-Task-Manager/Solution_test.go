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
		inputs [][]int
		ops    []op
		expect []int
	}{
		{"TestCase1", [][]int{
			{1, 101, 10}, {2, 102, 20}, {3, 103, 15},
		}, []op{
			{name: "add", userId: 4, taskId: 104, priority: 5},
			{name: "edit", taskId: 102, priority: 8},
			{name: "exec"},
			{name: "rmv", taskId: 101},
			{name: "add", userId: 5, taskId: 105, priority: 15},
			{name: "exec"},
		}, []int{3, 5}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.ops)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.ops)
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
