package Solution

import (
	"bytes"
	"container/list"
)

type TextEditor struct {
	cur  *list.Element
	list *list.List
}

func Constructor2296() TextEditor {
	return TextEditor{
		cur:  nil,
		list: list.New(),
	}
}

func (this *TextEditor) AddText(text string) {
	idx := 0
	if this.cur == nil {
		this.cur = this.list.PushFront(text[0])
		idx++
	}
	for ; idx < len(text); idx++ {
		v := this.list.InsertAfter(text[idx], this.cur)
		this.cur = v
	}
}

func (this *TextEditor) DeleteText(k int) int {
	if this.cur == nil {
		return 0
	}
	e := this.cur
	n := k
	for e != nil && k > 0 {
		pre := e.Prev()
		this.list.Remove(e)
		e = pre
		k--
	}
	this.cur = e
	return n - k
}

func (this *TextEditor) CursorLeft(k int) string {
	if this.cur == nil {
		return ""
	}
	for ; this.cur != nil && k > 0; this.cur, k = this.cur.Prev(), k-1 {
	}
	return this.pickup(10)
}

func (this *TextEditor) CursorRight(k int) string {
	if this.cur == nil && this.list.Len() == 0 {
		return ""
	}
	if this.cur == nil {
		this.cur = this.list.Front()
		k--
	}
	for ; this.cur.Next() != nil && k > 0; this.cur, k = this.cur.Next(), k-1 {
	}
	return this.pickup(10)
}

func (this *TextEditor) pickup(k int) string {
	if this.cur == nil {
		return ""
	}
	e := this.cur
	buf := bytes.NewBuffer([]byte{})
	for ; e != nil && k > 0; e, k = e.Prev(), k-1 {
		buf.WriteByte(e.Value.(byte))
	}
	bs := buf.Bytes()
	for s, e := 0, len(bs)-1; s < e; s, e = s+1, e-1 {
		bs[s], bs[e] = bs[e], bs[s]
	}
	return string(bs)
}

type opt struct {
	name string
	intOrStr
}

type intOrStr struct {
	a int
	b string
}

func Solution(opts []opt) []intOrStr {
	c := Constructor2296()
	ans := make([]intOrStr, 0)
	for _, op := range opts {
		if op.name == "a" {
			c.AddText(op.b)
			continue
		}
		if op.name == "d" {
			ans = append(ans, intOrStr{a: c.DeleteText(op.a)})
			continue
		}

		if op.name == "l" {
			ans = append(ans, intOrStr{b: c.CursorLeft(op.a)})
			continue
		}
		if op.name == "r" {
			ans = append(ans, intOrStr{b: c.CursorRight(op.a)})
		}
	}
	return ans
}
