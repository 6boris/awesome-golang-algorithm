package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                string
		queries, dictionary []string
		expect              []string
	}{
		{"TestCase1", []string{"word", "note", "ants", "wood"}, []string{"wood", "joke", "moat"}, []string{"word", "note", "wood"}},
		{"TestCase2", []string{"yes"}, []string{"not"}, nil},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.queries, c.dictionary)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.queries, c.dictionary)
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
