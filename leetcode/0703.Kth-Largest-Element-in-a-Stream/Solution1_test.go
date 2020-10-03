package Solution

import (
	"fmt"
	"testing"
)

func TestSolution1(t *testing.T) {
	kthLarges := Constructor1(3, []int{4, 5, 8, 2})

	fmt.Println(kthLarges.Add(3))
	fmt.Println(kthLarges.Add(5))
	fmt.Println(kthLarges.Add(10))
	fmt.Println(kthLarges.Add(9))
	fmt.Println(kthLarges.Add(4))
}

func TestSolution2(t *testing.T) {
	kthLarges := Constructor1(1, []int{})

	fmt.Println(kthLarges.Add(-3))
	fmt.Println(kthLarges.Add(-2))
	fmt.Println(kthLarges.Add(-4))
	fmt.Println(kthLarges.Add(0))
	fmt.Println(kthLarges.Add(4))
}
