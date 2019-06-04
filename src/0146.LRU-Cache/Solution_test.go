package Solution

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs int
		expect bool
	}{
		{"TestCase", 1, true},
		{"TestCase", 1, true},
		{"TestCase", 1, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Constructor(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				//t.Fatalf("expected: %v, but got: %v, with inputs: %v",
				//	c.expect, got, c.inputs)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	lru := Constructor(2)
	fmt.Println(lru.Get(2))

	lru.Put(2, 6)
	fmt.Println(lru.Get(1))

	lru.Put(1, 5)
	lru.Put(1, 2)
	lru.Print()

	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))

}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
