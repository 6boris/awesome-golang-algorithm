package Solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {

	kthLarges := Constructor(3, []int{4, 5, 8, 2})

	fmt.Println(kthLarges.Add(3))
	fmt.Println(kthLarges.Add(5))
	fmt.Println(kthLarges.Add(10))
	fmt.Println(kthLarges.Add(9))
	fmt.Println(kthLarges.Add(4))
}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
