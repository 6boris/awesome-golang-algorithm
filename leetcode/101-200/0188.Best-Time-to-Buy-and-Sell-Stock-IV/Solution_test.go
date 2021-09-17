package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution1(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect int
	}{
		{"TestCase", [][]int{
			{2, 4, 1},
			{2},
		}, 2},
		{"TestCase", [][]int{
			{3, 2, 6, 5, 0, 3},
			{7},
		}, 7},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+strconv.Itoa(i), func(t *testing.T) {
			got := maxProfit(c.inputs[0][0], c.inputs[0])
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
