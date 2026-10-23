package Solution

type Bitset struct {
	bits      []uint64
	size, one int

	flipped uint64
}

func Constructor(size int) Bitset {
	return Bitset{
		bits: make([]uint64, (size+63)/64),
		size: size,
	}
}

func (this *Bitset) Fix(idx int) {
	i, m := idx>>6, uint64(1)<<(idx&63)
	if (this.bits[i]^this.flipped)&m == 0 {
		this.bits[i] ^= m
		this.one++
	}
}

func (this *Bitset) Unfix(idx int) {
	i, m := idx>>6, uint64(1)<<(idx&63)
	if (this.bits[i]^this.flipped)&m != 0 {
		this.bits[i] ^= m
		this.one--
	}
}

func (this *Bitset) Flip() {
	this.flipped = ^this.flipped
	this.one = this.size - this.one
}

func (this *Bitset) All() bool  { return this.one == this.size }
func (this *Bitset) One() bool  { return this.one > 0 }
func (this *Bitset) Count() int { return this.one }

func (this *Bitset) ToString() string {
	buf := make([]byte, this.size)
	for i := range buf {
		w := this.bits[i>>6] ^ this.flipped
		buf[i] = '0' + byte(w>>(i&63)&1)
	}
	return string(buf)
}

type operation struct {
	name string
	idx  int
}

func Solution(size int, inputs []operation) []any {
	ret := make([]any, 0)
	c := Constructor(size)
	for i := range inputs {
		switch inputs[i].name {
		case "fix":
			c.Fix(inputs[i].idx)
		case "flip":
			c.Flip()
		case "unfix":
			c.Unfix(inputs[i].idx)
		case "all":
			ret = append(ret, c.All())
		case "one":
			ret = append(ret, c.One())
		case "count":
			ret = append(ret, c.Count())
		case "str":
			ret = append(ret, c.ToString())
		}
	}
	return ret
}
