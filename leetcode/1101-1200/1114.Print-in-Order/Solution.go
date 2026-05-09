package Solution

import (
	"bytes"
	"sync"
)

type Foo struct {
	b, c chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		b: make(chan struct{}, 1),
		c: make(chan struct{}, 1),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	f.b <- struct{}{}
}

func (f *Foo) Second(printSecond func()) {
	/// Do not change this line
	<-f.b
	printSecond()
	f.c <- struct{}{}
}

func (f *Foo) Third(printThird func()) {
	// Do not change this line
	<-f.c
	printThird()
}

func print1(buf *bytes.Buffer) func() {
	return func() {
		buf.WriteString("first")
	}
}

func print2(buf *bytes.Buffer) func() {
	return func() {
		buf.WriteString("second")
	}
}

func print3(buf *bytes.Buffer) func() {
	return func() {
		buf.WriteString("third")
	}
}

func Solution(nums []int) string {
	var buf bytes.Buffer
	foo := NewFoo()

	printFn := [3]func(){
		print1(&buf), print2(&buf), print3(&buf),
	}
	fooFn := [3]func(func()){
		foo.First, foo.Second, foo.Third,
	}
	var wg sync.WaitGroup
	wg.Add(3)
	for i := range nums {
		go func(i int) {
			defer wg.Done()
			fooFn[i](printFn[i])
		}(i)
	}
	wg.Wait()
	return buf.String()
}
