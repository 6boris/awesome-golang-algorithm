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
		arr1   []int
		arr2   []int
		d      int
		expect int
	}{
		{"TestCase", []int{4, 5, 8}, []int{10, 9, 1, 8}, 2, 2},
		{"TestCase", []int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3, 2},
		{"TestCase", []int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6, 1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.arr1, c.arr2, c.d)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.arr1, c.arr2, c.d)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
