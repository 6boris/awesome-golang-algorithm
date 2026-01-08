package Solution

import (
	"bytes"
)

type H2O struct {
	h, o chan struct{}
}

func NewH2O() *H2O {
	h := &H2O{
		h: make(chan struct{}, 2),
		o: make(chan struct{}, 1),
	}
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.h <- struct{}{}
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	h.o <- struct{}{}
	// releaseOxygen() outputs "H". Do not change or remove this line.
	releaseOxygen()
	<-h.h
	<-h.h
	<-h.o
}

func Solution(water string) string {
	h := make(chan struct{}, 10)
	o := make(chan struct{}, 10)
	buf := bytes.NewBuffer([]byte{})
	var releaseHydrogen, releaseOxygen func()
	releaseHydrogen = func() {
		buf.WriteByte('H')
	}
	releaseOxygen = func() {
		buf.WriteByte('O')
	}
	hh := NewH2O()
	done := make(chan struct{}, 2)
	go func() {
		defer func() {
			done <- struct{}{}
		}()
		for range h {
			hh.Hydrogen(releaseHydrogen)
		}
	}()
	go func() {
		defer func() {
			done <- struct{}{}
		}()

		for range o {
			hh.Oxygen(releaseOxygen)
		}
	}()

	for _, b := range []byte(water) {
		if b == 'H' {
			h <- struct{}{}
		}
		if b == 'O' {
			o <- struct{}{}
		}
	}

	close(h)
	close(o)
	for range 2 {
		<-done
	}

	return buf.String()
}
