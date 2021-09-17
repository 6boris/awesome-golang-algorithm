package Solution

import (
	"testing"
)

var N = 10

type Case struct {
	Name   string
	Input  *ListNode
	Expect *ListNode
}

func TestSolution1(t *testing.T) {
	//	测试用例
	cases := []Case{}
	N := 10
	for i := 0; i < N; i++ {
		name := "TestCacse1"
		input := UnmarshalListByRand(100, 10000)
		exect := Solution(input)
		cases = append(cases, Case{Name: name, Input: input, Expect: exect})
	}
	//	开始测试
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := Solution(c.Input)
			if isEqual(got, c.Expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.Expect, got, c.Input)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []Case{}
	N := 10
	for i := 0; i < N; i++ {
		name := "TestCacse2"
		input := UnmarshalListByRand(100, 100000)
		exect := Solution(input)
		cases = append(cases, Case{Name: name, Input: input, Expect: exect})
	}
	//	开始测试
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := Solution(c.Input)
			if isEqual(got, c.Expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.Expect, got, c.Input)
			}
		})
	}
}

func BenchmarkSolution1(b *testing.B) {
	//	测试用例
	cases := []Case{}
	N := 10
	for i := 0; i < N; i++ {
		name := "TestCacse1"
		input := UnmarshalListByRand(100, 1000000)
		exect := Solution(input)
		cases = append(cases, Case{Name: name, Input: input, Expect: exect})
	}
	//	开始测试
	for _, c := range cases {
		b.Run(c.Name, func(b *testing.B) {
			got := reverseList1(c.Input)
			if isEqual(got, c.Expect) {
				b.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.Expect, got, c.Input)
			}
		})
	}
}

func BenchmarkSolution2(b *testing.B) {
	//	测试用例
	cases := []Case{}
	N := 10
	for i := 0; i < N; i++ {
		name := "TestCacse1"
		input := UnmarshalListByRand(100, 100000)
		exect := Solution(input)
		cases = append(cases, Case{Name: name, Input: input, Expect: exect})
	}
	//	开始测试
	for _, c := range cases {
		b.Run(c.Name, func(b *testing.B) {
			got := reverseList2(c.Input)
			if isEqual(got, c.Expect) {
				b.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.Expect, got, c.Input)
			}
		})
	}
}
