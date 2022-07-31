package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name       string
		inputs     []int
		operations []string
		optNums    [][]int
		expect     []int
	}{
		{"TestCase1",
			[]int{1, 3, 5},
			[]string{"sumRange", "update", "sumRange"},
			[][]int{{0, 2}, {1, 2}, {0, 2}}, []int{9, 8}},
		{"TestCase2", []int{1},
			[]string{"sumRange", "update", "sumRange"},
			[][]int{{0, 0}, {0, 3}, {0, 0}}, []int{1, 3}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.operations, c.optNums)
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
