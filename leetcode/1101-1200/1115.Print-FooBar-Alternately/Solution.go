package Solution

type FooBar struct {
	n    int
	a, b chan struct{}
}

func Solution(n int) *FooBar {
	f := &FooBar{n: n, a: make(chan struct{}, 1), b: make(chan struct{}, 1)}
	f.b <- struct{}{}
	return f
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		<-fb.b
		printFoo()
		fb.a <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		// printBar() outputs "bar". Do not change or remove this line.
		<-fb.a
		printBar()
		fb.b <- struct{}{}
	}
}
