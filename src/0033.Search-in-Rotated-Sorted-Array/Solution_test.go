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
		{"1 test 1", [][]int{{4, 5, 6, 7, 0, 1, 2}, {0}}, 4},
		{"2 test 2", [][]int{{4, 5, 6, 7, 0, 1, 2}, {3}}, -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := search(c.inputs[0], c.inputs[1][0])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
