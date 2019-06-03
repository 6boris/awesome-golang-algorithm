package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect []int
	}{
		{
			"TestCase 1",
			[]int{1,2,3,3,4,4,5},
			[]int{1,2,5},
		},
		{
			"TestCase 2",
			[]int{1,1,1,2,3},
			[]int{2,3},
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Solution(UnmarshalListBySlice(c.inputs))
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
