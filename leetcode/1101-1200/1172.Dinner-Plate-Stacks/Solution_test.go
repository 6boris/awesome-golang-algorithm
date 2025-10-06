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
		inputs int
		opts   []opt
		expect []int
	}{
		{"TestCase1", 2, []opt{
			{"push", 1},
			{"push", 2},
			{"push", 3},
			{"push", 4},
			{"push", 5},
			{"popAtStack", 0},
			{"push", 20},
			{"push", 21},
			{"popAtStack", 0},
			{"popAtStack", 2},
			{"pop", 0},
			{"pop", 0},
			{"pop", 0},
			{"pop", 0},
			{"pop", 0},
		}, []int{2, 20, 21, 5, 4, 3, 1, -1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.opts)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.opts)
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
