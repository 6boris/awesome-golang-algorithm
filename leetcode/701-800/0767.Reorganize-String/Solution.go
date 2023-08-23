package Solution


import (
	"bytes"
	"fmt"
	"sort"
)

type tmp767 struct {
	b byte
	c int
}

func Solution(s string) string {

	list := make([]tmp767, 0)
	letterIndex := make(map[uint8]int)
	for _, b := range s {
		index, ok := letterIndex[byte(b)]
		if !ok {
			letterIndex[byte(b)] = len(list)
			list = append(list, tmp767{b: byte(b), c: 1})
			continue
		}
		list[index].c++
	}
	if len(list) == 1 {
		if list[0].c == 1 {
			return fmt.Sprintf("%c", list[0].b)
		}
		return ""
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].c == list[j].c {
			return list[i].b < list[j].b
		}
		return list[i].c < list[j].c
	})

	buf := bytes.NewBuffer([]byte{})
	count := len(list)
	for idx := 0; idx < count-1; idx++ {
		loop := list[idx].c
		for ; loop > 0; loop-- {
			for i := idx; i < len(list); i++ {
				buf.WriteByte(list[i].b)
				list[i].c--
			}
		}
	}
	bs := buf.Bytes()
	left := list[count-1].c
	if left == 0 {
		return string(bs)
	}
	bb := list[count-1].b
	bsl := len(bs)

	buf1 := bytes.NewBuffer([]byte{})
	if bb != bs[0] {
		buf1.WriteByte(bb)
		left--
	}

	for idx := 0; idx < bsl-1; idx++ {
		buf1.WriteByte(bs[idx])
		if bs[idx] != bb && bs[idx+1] != bb && left > 0 {
			buf1.WriteByte(bb)
			left--
		}
	}
	// write last byte
	buf1.WriteByte(bs[bsl-1])
	if left == 1 && bb != bs[bsl-1] {
		buf1.WriteByte(bb)
		left--
	}
	if left <= 0 {
		return buf1.String()
	}

	return ""
}
