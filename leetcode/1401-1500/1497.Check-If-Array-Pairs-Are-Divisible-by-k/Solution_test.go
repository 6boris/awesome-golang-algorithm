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
		k      int
		expect bool
	}{
		{"TestCase1", []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5, true},
		{"TestCase2", []int{1, 2, 3, 4, 5, 6}, 7, true},
		{"TestCase3", []int{1, 2, 3, 4, 5, 6}, 10, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.arr, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.arr, c.k)
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
