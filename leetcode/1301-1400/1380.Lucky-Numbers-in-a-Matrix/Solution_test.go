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
		expect []int
	}{
		{"TestCase", [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}, []int{15}},
		{"TestCase", [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, []int{12}},
		{"TestCase", [][]int{{7, 8}, {1, 2}}, []int{7}},
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
