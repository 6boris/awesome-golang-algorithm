package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name          string
		nums          []int
		target, start int
		expect        int
	}{
		{"TestCase1", []int{1, 2, 3, 4, 5}, 5, 3, 1},
		{"TestCase2", []int{1}, 1, 0, 0},
		{"TestCase3", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.target, c.start)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.nums, c.target, c.start)
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
