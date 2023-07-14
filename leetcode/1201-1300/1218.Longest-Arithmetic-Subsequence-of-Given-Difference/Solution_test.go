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
		arr    []int
		diff   int
		expect int
	}{
		{"TestCase1", []int{1, 2, 3, 4}, 1, 4},
		{"TestCase2", []int{1, 3, 5, 7}, 1, 1},
		{"TestCase3", []int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2, 4},
		{"TestCase4", []int{7, 7, 7, 7, 7, 7, 7}, 0, 7},
		{"TestCase5", []int{4, 12, 10, 0, -2, 7, -8, 9, -9, -12, -12, 8, 8}, 0, 2},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.arr, c.diff)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.arr, c.diff)
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
