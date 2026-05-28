package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                       string
		wordsContainer, wordsQuery []string
		expect                     []int
	}{
		{"TestCase1", []string{"abcd", "bcd", "xbcd"}, []string{"cd", "bcd", "xyz"}, []int{1, 1, 1}},
		{"TestCase2", []string{"abcdefgh", "poiuygh", "ghghgh"}, []string{"gh", "acbfgh", "acbfegh"}, []int{2, 0, 2}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.wordsContainer, c.wordsQuery)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.wordsContainer, c.wordsQuery)
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
