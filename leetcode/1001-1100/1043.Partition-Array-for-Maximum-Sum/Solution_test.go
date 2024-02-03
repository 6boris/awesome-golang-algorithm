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
		expect int
	}{
		{"TestCase1", []int{1, 15, 7, 9, 2, 5, 10}, 3, 84},
		{"TestCase2", []int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4, 83},
		{"TestCase3", []int{1}, 1, 1},
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
