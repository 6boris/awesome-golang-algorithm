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
		arr    []string
		k      int
		expect string
	}{
		{"TestCase1", []string{"d", "b", "c", "b", "c", "a"}, 2, "a"},
		{"TestCase2", []string{"aaa", "aa", "a"}, 1, "aaa"},
		{"TestCase3", []string{"a", "b", "a"}, 3, ""},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.arr, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.arr, c.k)
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
