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
		inputs []*NestedInteger
		expect []int
	}{
		{"TestCase1", []*NestedInteger{{Val: 1, isInt: true}, {isInt: false, list: []*NestedInteger{{Val: 2, isInt: true}, {Val: 3, isInt: true}}}, {Val: 1, isInt: true, list: nil}}, []int{1, 2, 3, 1}},
		{"TestCase2", []*NestedInteger{{Val: 1, isInt: true}, {Val: 2, isInt: true}, {Val: 3, isInt: true}}, []int{1, 2, 3}},
		{
			"TestCase3",
			[]*NestedInteger{
				{Val: 1, isInt: true},
				{isInt: false, list: []*NestedInteger{
					{isInt: false, list: []*NestedInteger{{Val: 2, isInt: true}}},
					{isInt: false, list: []*NestedInteger{
						{Val: 3, isInt: true},
						{Val: 4, isInt: true},
					}},
				}},
				{Val: 5, isInt: true},
			},
			[]int{1, 2, 3, 4, 5},
		},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
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
