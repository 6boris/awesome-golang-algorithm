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
		n           int
		left, right []int
		expect      bool
	}{
		{"TestCase1", 4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}, true},
		{"TestCase2", 4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1}, false},
		{"TestCase3", 2, []int{1, 0}, []int{-1, -1}, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.left, c.right)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n, c.left, c.right)
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
