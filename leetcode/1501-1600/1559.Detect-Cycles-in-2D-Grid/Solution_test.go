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
		inputs [][]byte
		expect bool
	}{
		{"TestCase1", [][]byte{
			{'a', 'a', 'a', 'a'},
			{'a', 'b', 'b', 'a'},
			{'a', 'b', 'b', 'a'},
			{'a', 'a', 'a', 'a'},
		}, true},
		{"TestCase2", [][]byte{
			{'c', 'c', 'c', 'a'}, {'c', 'd', 'c', 'c'},
			{'c', 'c', 'e', 'c'}, {'f', 'c', 'c', 'c'},
		}, true},
		{"TestCase1", [][]byte{
			{'a', 'b', 'b'}, {'b', 'z', 'b'}, {'b', 'b', 'a'},
		}, false},
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
