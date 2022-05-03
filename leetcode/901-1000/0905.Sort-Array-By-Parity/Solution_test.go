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
		expect []int
	}{
		{"TestCase1", []int{3, 1, 2, 4}, []int{2, 4, 1, 3}},
		{"TestCase2", []int{1, 3, 5, 7, 9, 2, 4, 8, 11}, []int{2, 4, 8, 1, 3, 5, 7, 9, 11}},
		{"TestCase3", []int{1}, []int{1}},
		{"testCase4", []int{2}, []int{2}},
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
		expect []int
	}{
		{"TestCase1", []int{3, 1, 2, 4}, []int{4, 2, 1, 3}},
		{"TestCase2", []int{1, 3, 5, 7, 9, 2, 4, 8, 11}, []int{8, 4, 2, 7, 9, 5, 3, 1, 11}},
		{"TestCase3", []int{1}, []int{1}},
		{"testCase4", []int{2}, []int{2}},
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

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
