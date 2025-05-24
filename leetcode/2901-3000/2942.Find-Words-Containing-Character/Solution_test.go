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
		inputs []string
		x      byte
		expect []int
	}{
		{"TestCase1", []string{"leet", "code"}, 'e', []int{0, 1}},
		{"TestCase2", []string{"abc", "bcd", "aaaa", "cbc"}, 'a', []int{0, 2}},
		{"TestCase3", []string{"abc", "bcd", "aaaa", "cbc"}, 'z', nil},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.x)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.x)
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
