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
		input1 string
		input2 string
		expect bool
	}{
		{"TestCase", "a", "b", false},
		{"TestCase", "aa", "ab", false},
		{"TestCase", "aa", "aab", true},
		{"TestCase", "aab", "baa", true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := canConstruct(c.input1, c.input2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with input1: %v input2: %v",
					c.expect, got, c.input1, c.input2)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 string
		input2 string
		expect bool
	}{
		{"TestCase", "a", "b", false},
		{"TestCase", "aa", "ab", false},
		{"TestCase", "aa", "aab", true},
		{"TestCase", "aab", "baa", true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := canConstruct2(c.input1, c.input2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with input1: %v input2: %v",
					c.expect, got, c.input1, c.input2)
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
