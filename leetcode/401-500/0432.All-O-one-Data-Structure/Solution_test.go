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
		expect []string
	}{
		{"TestCase1", []opt{
			{"i", "hello"},
			{"i", "hello"},
			{"max", ""},
			{"min", ""},
			{"i", "leet"},
			{"max", ""},
			{"min", ""},
		}, []string{"hello", "hello", "hello", "leet"}},
		{"TestCase2", []opt{
			{"i", "a"},
			{"i", "b"},
			{"i", "b"},
			{"i", "b"},
			{"i", "b"},
			{"d", "b"},
			{"d", "b"},
			{"max", ""},
			{"min", ""},
		}, []string{"b", "a"}},
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
