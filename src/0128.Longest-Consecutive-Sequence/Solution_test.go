package Soluation

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	// 测试用例
	cases := []struct {
		name 	string
		inputs	[]int
		expect 	int
	}{
		{
			"TestCases 1",
			[]int{100, 4, 200, 1, 3, 2},
			4,
		},
	}

	// 开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := longestConsecutive(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, got, c.inputs)
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
