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
		expect int
	}{
		{"TestCase1", [][]int{{1, 2}, {3, 4}}, 34},
		{"TestCase2", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, 32},
		{"TestCase3", [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}, 46},
		{"TestCase4", [][]int{{1}}, 6},
		{"TestCase5", [][]int{{2}}, 10},
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
