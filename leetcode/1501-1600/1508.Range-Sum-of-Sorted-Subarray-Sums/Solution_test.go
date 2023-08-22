package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name           string
		nums           []int
		n, left, right int
		expect         int
	}{
		{"TestCase1", []int{1, 2, 3, 4}, 4, 1, 5, 13},
		{"TestCase2", []int{1, 2, 3, 4}, 4, 3, 4, 6},
		{"TestCase3", []int{1, 2, 3, 4}, 4, 1, 10, 50},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.n, c.left, c.right)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.nums, c.n, c.left, c.right)
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
