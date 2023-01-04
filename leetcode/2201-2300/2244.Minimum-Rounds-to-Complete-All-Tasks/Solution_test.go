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
		expect int
	}{
		{"TestCase1", []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}, 4},
		{"TestCase2", []int{2, 3, 3}, -1},
		{"TestCase3", []int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7}, 12},
		{"TestCase4", []int{119, 115, 115, 119, 118, 113, 118, 120, 110, 113, 119, 115, 116, 118, 120, 117, 116, 111, 113, 119, 115, 113, 115, 111, 112, 119, 111, 111, 110, 112, 113, 120, 110, 111, 112, 111, 119, 112, 113, 112, 115, 116, 113, 114, 118, 119, 115, 114, 114, 112, 110, 117, 120, 110, 117, 116, 120, 118, 110, 120, 119, 113, 119, 120, 113, 110, 120, 114, 119, 115, 119, 117, 120, 116, 113, 113, 110, 118, 117, 116, 114, 114, 111, 116, 119, 112, 113, 116, 112, 116, 119, 112, 114, 114, 112, 118, 116, 113, 117, 116}, 38},
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

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
