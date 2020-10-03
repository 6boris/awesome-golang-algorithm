package Solution

import (
	"reflect"
	"runtime"
	"testing"
)

func TestSolution(t *testing.T) {
	if runtime.GOOS == "windows" {
		return
	}
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect string
	}{
		{"TestCacse 1", "/home/", "/home"},
		{"TestCacse 2", "/../", "/"},
		{"TestCacse 3", "/home//foo/", "/home/foo"},
		{"TestCacse 4", "/a/./b/../../c/", "/c"},
		{"TestCacse 5", "/a/../../b/../c//.//", "/c"},
		{"TestCacse 6", "/a//b////c/d//././/..", "/a/b/c"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := simplifyPath(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}
func TestSolution2(t *testing.T) {
	if runtime.GOOS == "windows" {
		return
	}

	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect string
	}{
		{"TestCacse 1", "/home/", "/home"},
		{"TestCacse 2", "/../", "/"},
		{"TestCacse 3", "/home//foo/", "/home/foo"},
		{"TestCacse 4", "/a/./b/../../c/", "/c"},
		{"TestCacse 5", "/a/../../b/../c//.//", "/c"},
		{"TestCacse 6", "/a//b////c/d//././/..", "/a/b/c"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := simplifyPath2(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
