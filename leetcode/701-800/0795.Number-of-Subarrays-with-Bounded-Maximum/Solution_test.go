package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name        string
		nums        []int
		left, right int
		expect      int
	}{
		{"TestCase1", []int{2, 1, 4, 3}, 2, 3, 3},
		{"TestCase2", []int{2, 9, 2, 5, 6}, 2, 8, 7},
		{"TestCase3", []int{1, 2, 2, 1}, 2, 3, 8},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.left, c.right)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.nums, c.left, c.right)
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
