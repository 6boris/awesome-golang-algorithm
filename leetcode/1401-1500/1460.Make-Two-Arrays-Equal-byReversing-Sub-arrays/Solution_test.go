package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name       string
		input, arr []int
		expect     bool
	}{
		{"TestCase1", []int{1, 2, 3, 4}, []int{2, 4, 1, 3}, true},
		{"TestCase2", []int{7}, []int{7}, true},
		{"TestCase3", []int{3, 7, 9}, []int{3, 7, 11}, false},
		{"TestCase4", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.input, c.arr)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.input, c.arr)
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
