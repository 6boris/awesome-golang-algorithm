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
		inputs [][]int
		cap    int
		expect bool
	}{
		{"TestCase1", [][]int{
			{2, 1, 5},
			{3, 3, 7},
		}, 4, false},
		{"TestCase2", [][]int{
			{2, 1, 5},
			{3, 3, 7},
		}, 5, true},
		{"TestCase3", [][]int{
			{1, 1, 8},
			{2, 2, 7},
			{3, 3, 6},
		}, 5, false},
		{"TestCase4", [][]int{
			{1, 1, 8},
			{2, 2, 7},
			{3, 3, 6},
		}, 6, true,
		},
		{"TestCase5", [][]int{
			{2, 1, 5},
			{3, 5, 7},
		}, 3, true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.cap)
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
