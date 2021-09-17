package Solution

import (
	"testing"
)

type InputCase struct {
	element []int
	k       int
}

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs InputCase
		expect []int
	}{
		{
			"TestCase 1",
			InputCase{
				[]int{1, 2, 3, 4, 5},
				2,
			},
			[]int{4, 5, 1, 2, 3},
		},
		{
			"TestCase 2",
			InputCase{
				[]int{0, 1, 2},
				4,
			},
			[]int{2, 0, 1},
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			listNode := UnmarshalListBySlice(c.inputs.element)
			ret := Solution(listNode, c.inputs.k)
			if !isEqual(ret, UnmarshalListBySlice(c.expect)) {
				PrintList(ret)
				PrintList(UnmarshalListBySlice(c.expect))
				t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, ret, c.inputs)
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
