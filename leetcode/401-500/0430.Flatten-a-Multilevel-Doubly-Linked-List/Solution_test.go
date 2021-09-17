package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	t.Run("TestCase", func(t *testing.T) {
		main := MakeNode([]int{1, 2, 3, 4, 5, 6})
		child1 := MakeNode([]int{7, 8, 9, 10})
		child2 := MakeNode([]int{11, 12})
		child1.Next.Child = child2
		curr := main
		for i := 0; i < 2; i++ {
			curr = curr.Next
		}
		curr.Child = child1
		got := Solution(main)
		expect := MakeNode([]int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6})
		if !Check(got, expect) {
			t.Fatalf("expected: %v, but got: %v, with inputs: %v",
				expect, got, main)
		}
	})
	t.Run("TestCase", func(t *testing.T) {
		main := MakeNode([]int{1, 2})
		main.Child = &Node{Val: 3}
		got := Solution(main)
		expect := MakeNode([]int{1, 3, 2})
		if !Check(got, expect) {
			t.Fatalf("expected: %v, but got: %v, with inputs: %v",
				expect, got, main)
		}
	})
	t.Run("TestCase", func(t *testing.T) {
		main := MakeNode([]int{})
		got := Solution(main)
		expect := MakeNode([]int{})
		if !Check(got, expect) {
			t.Fatalf("expected: %v, but got: %v, with inputs: %v",
				expect, got, main)
		}
	})
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
