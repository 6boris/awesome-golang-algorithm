package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name      string
		inputs    [][]int
		k, expect int
	}{
		{"TestCase1", [][]int{{1, 5, 9}, {10, 11, 13, 12, 13, 15}}, 8, 13},
		{"TestCase2", [][]int{{-5}}, 1, -5},
		{"TestCase3", [][]int{{1}, {2}, {3}}, 3, 3},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.k)
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
