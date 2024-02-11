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
		expect      int
	}{
		{"TestCase1", 4, []int{4, 3}, []int{0, 1}, 4},
		{"TestCase2", 7, []int{}, []int{0, 1, 2, 3, 4, 5, 6, 7}, 7},
		{"TestCase3", 7, []int{0, 1, 2, 3, 4, 5, 6, 7}, []int{}, 7},
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
