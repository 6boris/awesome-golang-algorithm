package Solution

import (
	"fmt"
	"testing"
)

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	obj := Constructor()
	obj.Push_back(1)
	obj.Push_back(2)
	param_1 := obj.Max_value()
	param_2 := obj.Pop_front()
	param_3 := obj.Max_value()
	fmt.Println(param_1, param_2, param_3)
}

func TestSolution2(t *testing.T) {
	obj := Constructor()
	fmt.Println(obj.Pop_front())
	fmt.Println(obj.Pop_front())
	fmt.Println(obj.Pop_front())
	fmt.Println(obj.Pop_front())
	fmt.Println(obj.Pop_front())
	obj.Push_back(15)
	fmt.Println(obj.Max_value())
	obj.Push_back(9)
	fmt.Println(obj.Max_value())
}
