package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name     string
		bookings [][]int
		n        int
		expect   []int
	}{
		{"TestCase1", [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5, []int{10, 55, 45, 25, 25}},
		{"TestCase2", [][]int{{1, 2, 10}, {2, 2, 15}}, 2, []int{10, 25}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.bookings, c.n)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.bookings, c.n)
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
