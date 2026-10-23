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
		target int
		expect int
	}{
		{"TestCase1", []int{3, 2, 2, 4, 3}, 3, 2},
		{"TestCase2", []int{7, 3, 4, 7}, 7, 2},
		{"TestCase3", []int{4, 3, 2, 6, 2, 3, 4}, 6, -1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.arr, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.arr, c.target)
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
