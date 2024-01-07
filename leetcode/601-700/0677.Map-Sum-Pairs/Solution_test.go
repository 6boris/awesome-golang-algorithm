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
		inputs []opt
		expect []int
	}{
		{"TestCase1", []opt{{"i", "apple", 3}, {"", "ap", 0}, {"i", "app", 2}, {"", "ap", 0}}, []int{3, 5}},
		{"TestCase2", []opt{{"i", "apple", 3}, {"", "ap", 0}, {"i", "app", 2}, {"i", "apple", 2}, {"", "ap", 0}}, []int{3, 4}},
		{"TestCase3", []opt{{"i", "apple", 3}, {"", "app", 0}, {"i", "app", 2}, {"", "ap", 0}, {"i", "apple", 5}, {"i", "apple", 1}, {"", "apple", 0}}, []int{3, 5, 1}},
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

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
