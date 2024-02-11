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
		{"TestCase1", []string{"9001 discuss.leetcode.com"}, []string{"9001 com", "9001 discuss.leetcode.com", "9001 leetcode.com"}},
		{"TestCase2", []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}, []string{"1 intel.mail.com", "5 org", "5 wiki.org", "50 yahoo.com", "900 google.mail.com", "901 mail.com", "951 com"}},
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
