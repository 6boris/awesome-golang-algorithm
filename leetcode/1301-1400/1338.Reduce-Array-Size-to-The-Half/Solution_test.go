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
		expect int
	}{
		{"TestCase", []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}, 2},
		{"TestCase", []int{7, 7, 7, 7, 7, 7}, 1},
		{"TestCase", []int{1000, 1000, 3, 7}, 1},
		{"TestCase", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution1(c.nums)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.expect)
			}
		})
	}
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution2(c.nums)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.expect)
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
