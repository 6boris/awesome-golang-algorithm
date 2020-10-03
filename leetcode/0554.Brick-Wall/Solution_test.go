package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect int
	}{
		{"1 test 1", [][]int{
			{1, 2, 2, 1},
			{3, 1, 2},
			{1, 3, 2},
			{2, 4},
			{3, 1, 2},
			{1, 3, 1, 1},
		}, 2},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := leastBricks(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
