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
		inputs []int
		d      int
		expect int
	}{
		{"TestCase1", []int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2, 4},
		{"TestCase2", []int{3, 3, 3, 3, 3}, 3, 1},
		{"TestCase3", []int{7, 6, 5, 4, 3, 2, 1}, 1, 7},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.d)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.d)
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
