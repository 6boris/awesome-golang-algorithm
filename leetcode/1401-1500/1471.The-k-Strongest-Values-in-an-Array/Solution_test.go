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
		expect []int
	}{
		{"TestCase1", []int{1, 2, 3, 4, 5}, 2, []int{5, 1}},
		{"TestCase2", []int{1, 1, 3, 5, 5}, 2, []int{5, 5}},
		{"TestCase3", []int{6, 7, 11, 7, 6, 8}, 5, []int{11, 8, 6, 6, 7}},
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
