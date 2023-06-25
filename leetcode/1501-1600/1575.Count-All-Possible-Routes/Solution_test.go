package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                string
		locations           []int
		start, finish, fuel int
		expect              int
	}{
		{"TestCase1", []int{2, 3, 6, 8, 4}, 1, 3, 5, 4},
		{"TestCase2", []int{22, 74, 92, 86, 12, 68, 64, 19, 79, 10, 69, 13, 62, 18, 87, 88, 33, 96, 78, 73, 57, 42, 91, 17, 55, 26, 27, 67, 60, 46, 72, 41}, 30, 29, 47, 535415296},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.locations, c.start, c.finish, c.fuel)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.locations, c.start, c.finish, c.fuel)
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
