package Solution

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []string
		expect []string
	}{
		{"TestCase1", []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}, []string{"/a", "/c/d", "/c/f"}},
		{"TestCase2", []string{"/a", "/a/b/c", "/a/b/d"}, []string{"/a"}},
		{"TestCase3", []string{"/a/b/c", "/a/b/ca", "/a/b/d"}, []string{"/a/b/c", "/a/b/ca", "/a/b/d"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			sort.Strings(got)
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
