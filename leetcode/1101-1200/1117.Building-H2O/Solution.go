package Solution

import (
	"bytes"
)

type H2O struct {
	hChan chan struct{}
	oChan chan struct{}
	done  chan struct{}
}

func NewH2O() *H2O {
	return &H2O{
		hChan: make(chan struct{}, 2), // 允许 2 个 H 进入
		oChan: make(chan struct{}, 1), // 允许 1 个 O 进入
		done:  make(chan struct{}),    // 用于三者同步
	}
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.hChan <- struct{}{}

	releaseHydrogen()
	h.done <- struct{}{}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	h.oChan <- struct{}{}

	releaseOxygen()

	<-h.done
	<-h.done

	<-h.hChan
	<-h.hChan
	<-h.oChan
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
