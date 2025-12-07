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
		k      int
		expect int
	}{
		{"TestCase1", [][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}}, 3, 2},
		{"TestCase2", [][]int{{0, 0}}, 5, 1},
		{"TestCase3", [][]int{{7, 3, 4, 9}, {2, 3, 6, 2}, {2, 3, 7, 0}}, 1, 10},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.k)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
