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
		inputs [][]string
		expect [][]string
	}{
		{"TestCase1", [][]string{
			{"a"}, {"c"}, {"d"}, {"a", "b"},
			{"c", "b"}, {"d", "a"},
		}, [][]string{{"d"}, {"d", "a"}}},
		{"TestCase2", [][]string{
			{"a"}, {"c"}, {"a", "b"}, {"c", "b"}, {"a", "b", "x"},
			{"a", "b", "x", "y"}, {"w"}, {"w", "y"},
		}, [][]string{
			{"a"}, {"a", "b"},
			{"c"}, {"c", "b"},
		}},
		{"TestCase3", [][]string{
			{"a", "b"}, {"c", "d"}, {"c"}, {"a"},
		}, [][]string{
			{"a"}, {"a", "b"}, {"c"}, {"c", "d"},
		}},
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
