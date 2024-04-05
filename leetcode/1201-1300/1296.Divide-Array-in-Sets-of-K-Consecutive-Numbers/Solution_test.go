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
		nums   []int
		k      int
		expect bool
	}{
		{"TestCase1", []int{1, 2, 3, 3, 4, 4, 5, 6}, 4, true},
		{"TestCase2", []int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3, true},
		{"TestCase3", []int{1, 2, 3, 4}, 3, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.nums, c.k)
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
