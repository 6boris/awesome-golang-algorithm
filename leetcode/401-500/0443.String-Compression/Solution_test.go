package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		inputs  []byte
		expect1 []byte
		expect  int
	}{
		{"TestCase1", []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, []byte{'a', '2', 'b', '2', 'c', '3'}, 6},
		{"TestCase2", []byte{'a'}, []byte{'a'}, 1},
		{"TestCase3", []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, []byte{'a', 'b', '1', '2'}, 4},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
			if !reflect.DeepEqual(c.inputs[:got], c.expect1) {
				t.Fatalf("expect %v but got: %v", c.expect1, c.inputs[:got])
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
