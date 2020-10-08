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
		expect []int
	}{
		{"TestCase", []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}, []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
		{"TestCase", []int{2, 6, 2, 1, 5, 1, 4, 6}, []int{1, 2, 4}, []int{1, 1, 2, 2, 4, 5, 6, 6}},
		{"TestCase", []int{3, 3, 2, 2, 1, 1}, []int{}, []int{1, 1, 2, 2, 3, 3}},
		{"TestCase", []int{}, []int{}, []int{}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.arr1, c.arr2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.arr1, c.arr2)
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
