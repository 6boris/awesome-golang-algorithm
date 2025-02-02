package Solution

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"sync"
	"testing"
)

func foo(output io.Writer) func() {
	return func() {
		fmt.Fprint(output, "foo")
	}
}

func bar(output io.Writer) func() {
	return func() {
		fmt.Fprint(output, "bar")
	}
}
func run(fb *FooBar) string {
	var wg sync.WaitGroup
	wg.Add(2)
	buf := bytes.NewBuffer([]byte{})
	go func() {
		defer wg.Done()
		fb.Foo(foo(buf))
	}()
	go func() {
		defer wg.Done()
		fb.Bar(bar(buf))
	}()
	wg.Wait()
	return buf.String()
}

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs int
		expect string
	}{
		{"TestCase1", 1, "foobar"},
		{"TestCase2", 2, "foobarfoobar"},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := run(Solution(c.inputs))
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
