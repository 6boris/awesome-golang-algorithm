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
		inputs []int
		expect int64
	}{
		{"TestCase1", []int{3, 2, 5, 1, 3, 4}, 22},
		{"TestCase2", []int{3, 4}, 12},
		{"TestCase3", []int{1, 1, 2, 3}, -1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestSolution1(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int64
	}{
		{"TestCase1", []int{3, 2, 5, 1, 3, 4}, 22},
		{"TestCase2", []int{3, 4}, 12},
		{"TestCase3", []int{1, 1, 2, 3}, -1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution1(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
