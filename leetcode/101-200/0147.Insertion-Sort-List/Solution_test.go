package Solution

import "testing"

func TestSolution(t *testing.T) {
	TestCases := []struct {
		name   string
		inputs []int
		expect []int
	}{
		{
			"TestCase 1",
			[]int{4, 2, 1, 3},
			[]int{1, 2, 3, 4},
		},
		{
			"TestCase 2",
			[]int{-1, 5, 3, 4, 0},
			[]int{-1, 0, 3, 4, 5},
		},
	}

	for _, c := range TestCases {
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
