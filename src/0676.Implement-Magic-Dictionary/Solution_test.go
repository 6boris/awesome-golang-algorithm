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
		dict   []string
		inputs string
		expect bool
	}{
		{"TestCase", []string{"hello", "leetcode"}, "hello", false},
		{"TestCase", []string{"hello", "leetcode"}, "hhllo", true},
		{"TestCase", []string{"hello", "leetcode"}, "hell", false},
		{"TestCase", []string{"hello", "leetcode"}, "leetcoded", false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			obj := Constructor()
			obj.BuildDict(c.dict)
			got := obj.Search(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with dict: %v, search: %v",
					c.expect, got, c.dict, c.inputs)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
